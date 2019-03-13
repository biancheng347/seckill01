package config

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	 "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"seckill01/models"
	"time"
)

var (
	etcdClient *clientv3.Client
)

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func initLogger() (err error)  {
	config := make(map[string] interface{})
	config["filename"] = secKillConf.LogPath
	config["level"] = convertLogLevel(secKillConf.LogLevel)

	configByte,err := json.Marshal(config)
	if err != nil {
		err = fmt.Errorf("initLogger configByte failed,err:%v",err)
		return
	}

	logs.SetLogger(logs.AdapterFile,string(configByte))
	return
}

func updateSecProductInfo(secproductInfo []models.SecProductInfoConf) {
	tmp := make(map[int]*models.SecProductInfoConf,1024)
	for _,v := range secproductInfo {
		productInfo := v
		tmp[v.ProductId] = &productInfo
	}

	secKillConf.RWSecProductLock.Lock()
	secKillConf.SecProductInfoMap = tmp
	secKillConf.RWSecProductLock.Unlock()
}


func loadSecConfig() (err error) {
	go func() {
		resp,err := etcdClient.Get(context.Background(),secKillConf.EtcdConf.EtcdSecProductKey)
		if err != nil {
			logs.Error("get productKey failed: %v:%v",secKillConf.EtcdConf.EtcdSecProductKey,err)
			return
		}
		var secProductInfo []models.SecProductInfoConf
		for _,v := range resp.Kvs {
			if err = json.Unmarshal(v.Value,&secProductInfo); err != nil {
				logs.Error("unmarshal sec product key info failed,err: %v",err)
				return
			}
		}
		updateSecProductInfo(secProductInfo)
	}()
	return
}

func initEtcd() (err error) {
	cli,err := clientv3.New(clientv3.Config{
		Endpoints: []string{secKillConf.EtcdConf.EtcdAddr},
		DialTimeout: time.Duration(secKillConf.EtcdConf.Timeout) * time.Second,
	})
	if err != nil {
		logs.Error("connct etcd failed,err:",err)
		return
	}
	etcdClient = cli
	return
}



func watchSecProductKey(key string) {
	cli,err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed,err: %v",err)
		return
	}

	for {
		var secProductInfo []models.SecProductInfoConf
		getConsucc := true
		for watchResp := range  cli.Watch(context.Background(),key){
			for _,ev := range watchResp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key: %v, is deleted",key)
					continue
				}else if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					if err = json.Unmarshal(ev.Kv.Value,&secProductInfo); err != nil {
						logs.Error("key: %v failed,err: %v",key,err)
						getConsucc = false
						continue
					}
				}
			}
			if  getConsucc {
				updateSecProductInfo(secProductInfo)
			}
		}
	}
}


func initSecProcutWatcher() {
	go watchSecProductKey(secKillConf.EtcdConf.EtcdSecProductKey)
}

func InitSecKill() (err error) {
	if err = initLogger();err != nil {
		return
	}

	if err = initEtcd();err != nil {
		return
	}

	if err = loadSecConfig();err != nil {
		return
	}
	models.InitServer(secKillConf)
	initSecProcutWatcher()
	return
}