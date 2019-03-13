package structModel

import (
	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/clientv3"
	"fmt"
	"strings"
	"time"
)

type EtcdConfParam struct{
	EtcdAddr          string
	Timeout           int
}

type EtcdConf struct{
	EtcdConfParam
	EtcdSecKeyPrefix  string
	EtcdSecProductKey string
}


func (p *EtcdConf) InitEtcdConf(keys ...string) (err error) {
	for _, v := range keys {
		if strings.HasSuffix(v, "addr") {
			if err = appConfigStringValue(&p.EtcdAddr, v); err != nil {
				break
			}
		} else if strings.HasSuffix(v, "timeout") {
			if err = appConfigIntValue(&p.Timeout, v); err != nil {
				break
			}
		} else if strings.HasSuffix(v, "prefix") {
			if err = appConfigStringValue(&p.EtcdSecKeyPrefix, v); err != nil {
				break
			}
		} else if strings.HasSuffix(v, "key") {
			if err = appConfigStringValue(&p.EtcdSecProductKey, v); err != nil {
				break
			}
			if err = p.etcdproductKey(v); err != nil {
				break
			}
		}
	}
	return
}

func (p *EtcdConf) etcdproductKey(key string) (err error) {
	productKey := ""
	if err = appConfigStringValue(&productKey,key);err != nil {
		return
	}
	if strings.HasSuffix(p.EtcdSecKeyPrefix, "/") == false {
		p.EtcdSecKeyPrefix = p.EtcdSecKeyPrefix + "/"
	}
	p.EtcdSecProductKey = fmt.Sprintf("%s%s", p.EtcdSecKeyPrefix, productKey)
	return
}


func (p EtcdConfParam)InitEtcdConf() (cli *clientv3.Client, err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{p.EtcdAddr},
		DialTimeout: time.Duration(p.Timeout) * time.Second,
	})
	if err != nil {
		logs.Error("connct etcd failed,err:", err)
		return
	}
	return
}