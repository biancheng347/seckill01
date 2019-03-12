package models

import (
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

func loadBlackList() (err error) {
	seckillconf.ipBlackMap = make(map[string]bool, 10000)
	seckillconf.idBlackMap = make(map[int]bool, 10000)

	pool,err := initRedis(seckillconf.RedisBlackConf)
	if err != nil {
		logs.Error("init black redis failed,err: %v",err)
		return
	}
	seckillconf.BlackRedisPool = pool

	conn := seckillconf.BlackRedisPool.Get()
	defer  conn.Close()

	relply,err := conn.Do("hgetall","idblacklist")
	idlist,err := redis.Strings(relply,err)
	if err != nil {
		logs.Warn("hget all failed,err:%v",err)
		return
	}

	for _,v := range idlist {
		id,err := strconv.Atoi(v)
		if err != nil {
			logs.Warn("invalid user id: %v",id)
			continue
		}
		seckillconf.idBlackMap[id] = true
	}

	relply ,err = conn.Do("hgetall","ipblacklist")
	iplist,err := redis.Strings(relply,err)
	if err != nil {
		logs.Warn("hget all failed,err:%v",err)
		return
	}

	for _,v := range iplist {
		seckillconf.ipBlackMap[v] = true
	}
	return
}




func initProxyToLayerRedis() (err error) {
	pool,err := initRedis(seckillconf.RedisProxyToLayerConf)
	if err != nil {
		logs.Error("init black redis failed,err: %v",err)
		return
	}
	seckillconf.ProxyToLayerRedisPool = pool
	return
}

func WriteHandle() {

}

func ReadHandle() {

}

func initRedisProcessFunc() {
	for i := 0; i < seckillconf.WriteProxyToLayerGoroutineNum; i++ {
		go WriteHandle()
	}

	for i := 0; i < seckillconf.ReadProxyToLayerGoroutineNum; i++ {
		go ReadHandle()
	}
}

func InitServer(secKillConfig *SecKillConf) (err error) {
	seckillconf = secKillConfig

	if err = loadBlackList();err != nil {
		logs.Error("load black list err: %v",err)
		return
	}
	logs.Debug("init service success,config: %v",seckillconf)

	if err = initProxyToLayerRedis();err != nil {
		logs.Error("load proxy2layer redis pool failed, err:%v", err)
		return
	}

	seckillconf.secLimitMgr = &SecLimitMgr{
		UserLimitMap:make(map[int]*Limit,10000),
		IpLimitMap:make(map[string]*Limit,10000),
	}
	seckillconf.SecReqChan = make(chan *SecRequest,seckillconf.SecReqChanSize)
	seckillconf.UserConnMap = make(map[string]chan *SecResult,10000)

	initRedisProcessFunc()

	return
}
