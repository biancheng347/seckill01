package base

import (
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"strings"
	"time"
)

type RedisConf struct {
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
}

func (p *RedisConf) InitRedisConf(keys ...string) (err error) {
	for _, v := range keys {
		if strings.HasSuffix(v, "addr") {
			if err = appConfigStringValue(&p.RedisAddr, v); err != nil {
				break
			}
		} else if strings.HasSuffix(v, "idle") {
			if err = appConfigIntValue(&p.RedisMaxIdle, v); err != nil {
				break
			}
		} else if strings.HasSuffix(v, "active") {
			if err = appConfigIntValue(&p.RedisMaxActive, v); err != nil {
				break
			}
		} else if strings.HasSuffix(v, "timeout") {
			if err = appConfigIntValue(&p.RedisIdleTimeout, v); err != nil {
				break
			}
		}
	}
	return
}

func (p RedisConf) initRedis() (redisPool *redis.Pool, err error) {
	pool := &redis.Pool{
		MaxIdle:     p.RedisMaxIdle,
		MaxActive:   p.RedisMaxActive,
		IdleTimeout: time.Duration(p.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", p.RedisAddr)
		},
	}
	conn := pool.Get()
	defer conn.Close()

	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed,err :%v", err)
		return
	}
	redisPool = pool
	return
}

func (p RedisConf)InitRedisValue(redisPool **redis.Pool) (err error) {
	pool,err := p.initRedis()
	if err != nil {
		logs.Error("init redis failed,err: %v,addr: %v",err,p.RedisAddr)
		return
	}
	*redisPool = pool
	return
}

type SecProductInfoConf struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int
	Total     int
	Remain    int
}


