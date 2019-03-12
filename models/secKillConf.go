package models

import (
	"github.com/garyburd/redigo/redis"
	"sync"
)

type RedisConf struct {
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
}

type SecProductInfoConf struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int
	Total     int
	Remain    int
}

type AccessLimitConf struct {
	IPSecAccessLimit   int
	UserSecAccessLimit int
	IPMinAccessLimit   int
	UserMinAccessLimit int
}

type EtcdConf struct{
	EtcdAddr          string
	Timeout           int
	EtcdSecKeyPrefix  string
	EtcdSecProductKey string
}


type SecKillConf struct {
	RedisBlackConf        RedisConf
	RedisLayerToProxyConf RedisConf
	RedisProxyToLayerConf RedisConf

	BlackRedisPool        *redis.Pool
	ProxyToLayerRedisPool *redis.Pool
	LayerToProxyRedisPool *redis.Pool

	WriteLayerToProxyGoroutineNum int
	ReadLayerToProxyGoroutineNum  int
	WriteProxyToLayerGoroutineNum int
	ReadProxyToLayerGoroutineNum  int

	idBlackMap        map[int]bool //
	ipBlackMap        map[string]bool //

	LogPath  string
	LogLevel string

	RWSecProductLock  sync.RWMutex
	secProductInfoMap map[int]*SecProductInfoConf //

	SecReqChan     chan *SecRequest
	SecReqChanSize int

	EtcdConf EtcdConf
	CookieSecretKey string
	ReferWhiteList []string
	AccessLimitConf   AccessLimitConf
	secLimitMgr       *SecLimitMgr //
	UserConnMap     map[string]chan *SecResult //
}

func NewSecKillConf() *SecKillConf {
	return &SecKillConf{
		secProductInfoMap: make(map[int]*SecProductInfoConf, 1024),
		idBlackMap: make(map[int]bool,10000),
		ipBlackMap: make(map[string]bool,10000),
		secLimitMgr: &SecLimitMgr{
			UserLimitMap:make(map[int]*Limit,10000),
			IpLimitMap:make(map[string]*Limit,10000),
		},
		SecReqChan: make(chan *SecRequest,10000),
		UserConnMap: make(map[string]chan *SecResult,10000),
	}
}

