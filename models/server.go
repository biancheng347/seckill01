package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	seckillconf *SecKillConf
)

func initBlackRedis() (err error) {
	seckillconf.BlackRedisPool = &redis.Pool{
		MaxIdle:     seckillconf.RedisBlackConf.RedisMaxIdle,
		MaxActive:   seckillconf.RedisBlackConf.RedisMaxActive,
		IdleTimeout: time.Duration(seckillconf.RedisBlackConf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", seckillconf.RedisBlackConf.RedisAddr)
		},
	}

	conn := seckillconf.BlackRedisPool.Get()
	defer conn.Close()

	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping readis failed,err :%v", err)
		return
	}
	return
}

func loadBlackList() (err error) {
	seckillconf.ipBlackMap = make(map[string]bool, 10000)
	seckillconf.idBlackMap = make(map[int]bool, 10000)

	return
}

func InitServer(secKillConfig *SecKillConf) (err error) {
	seckillconf = secKillConfig

	return
}
