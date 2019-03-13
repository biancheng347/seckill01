package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

var (
	seckillconf *SecKillConf
)

func initRedis(conf RedisConf) (redisPool  *redis.Pool,err error) {
	pool := &redis.Pool{
		MaxIdle: conf.RedisMaxIdle,
		MaxActive: conf.RedisMaxActive,
		IdleTimeout: time.Duration(conf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",conf.RedisAddr)
		},
	}
	conn := pool.Get()
	defer  conn.Close()

	_,err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed,err :%v", err)
		return
	}
	redisPool = pool
	return
}

func connDo(conn redis.Conn,name,args string,f func(list []string)) (err error) {
	relply,err := conn.Do(name,args)
	list,err := redis.Strings(relply,err)
	if err != nil {
		logs.Warn("name: %v , args: %v  failed,err:%v",name,args,err)
		return
	}
	f(list)
	return
}

func loadBlackList() (err error) {
	if err = initRedisValue(&seckillconf.BlackRedisPool,seckillconf.RedisBlackConf);err != nil {
		return
	}
	conn := seckillconf.BlackRedisPool.Get()
	defer  conn.Close()

	err = connDo(conn,"hgetall","idblacklist", func(list []string) {
		for _,v := range list {
			id,err := strconv.Atoi(v)
			if err != nil {
				logs.Warn("invalid user id: %v",id)
				continue
			}
			seckillconf.idBlackMap[id] = true
		}
	})
	if err != nil {
		return
	}

	err = connDo(conn,"hgetall","ipblacklist", func(list []string) {
		for _,v := range list {
			seckillconf.ipBlackMap[v] = true
		}
	})
	if err != nil {
		return
	}
	return
}

func initRedisValue(redisPool **redis.Pool,conf RedisConf) (err error) {
	pool,err := initRedis(conf)
	if err != nil {
		logs.Error("init redis failed,err: %v,addr: %v",err,conf.RedisAddr)
		return
	}
	*redisPool = pool
	return
}


func initProxyToLayerRedis() (err error) {
	if err = initRedisValue(&seckillconf.ProxyToLayerRedisPool,seckillconf.RedisProxyToLayerConf);err != nil {
		return
	}
	return
}

func WriteHandle() {
	f := func(req *SecRequest) {
		conn := seckillconf.ProxyToLayerRedisPool.Get()
		defer  conn.Close()

		data,err := json.Marshal(req)
		if err != nil {
			logs.Error("json marshal failed,err:%v, req: %v",err,req)
			return
		}

		if _,err = conn.Do("LPUSH","sec_queue",string(data));err != nil {
			logs.Error("lpush failed,err:%v, req: %v",err,req)
			return
		}
		return
	}
	for  {
		req := <- seckillconf.SecReqChan
		f(req)
	}
}

func ReadHandle() {
	f := func () {
		conn := seckillconf.ProxyToLayerRedisPool.Get()
		defer conn.Close()

		replay,err := conn.Do("RPOP","recv_queueu")
		data,err := redis.String(replay,err)
		if err == redis.ErrNil {
			time.Sleep(time.Second)
			return
		}else if err != nil {
			logs.Error("rpop failed,err: %v",err)
			return
		}
		logs.Debug("rpop from redis succ: data: %s",string(data))
		var result SecResult
		if	err = json.Unmarshal([]byte(data),&result); err != nil {
			logs.Error("json unmarshal failed,err:%v",err)
			return
		}

		userkey := fmt.Sprintf("%s_%s",result.UserId,result.ProductId)

		seckillconf.UserConnMapLock.Lock()
		resultChan,ok := seckillconf.UserConnMap[userkey]
		seckillconf.UserConnMapLock.Unlock()
		if !ok {
			logs.Warn("user not found: %v",userkey)
			return
		}

		resultChan <- &result
		return
	}
	for {
		f()
	}
}

func initRedisProcessFunc() {
	for i := 0; i < seckillconf.WriteProxyToLayerGoroutineNum; i++ {
		go WriteHandle()
	}

	for i := 0; i < seckillconf.ReadProxyToLayerGoroutineNum; i++ {
		go ReadHandle()
	}
}

func InitServer(sec *SecKillConf) (err error) {
	seckillconf = sec

	if err = loadBlackList();err != nil {
		logs.Error("load black list err: %v",err)
		return
	}
	logs.Debug("init service success,config: %v",seckillconf)

	if err = initProxyToLayerRedis();err != nil {
		logs.Error("load proxy2layer redis pool failed, err:%v", err)
		return
	}

	initRedisProcessFunc()
	return
}
