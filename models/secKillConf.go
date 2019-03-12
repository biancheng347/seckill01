package models

import "sync"

var (
	seckillconf = NewSecKillConf()
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

type SecKillConf struct {
	RedisBlackConf        RedisConf
	RedisLayerToProxyConf RedisConf
	RedisProxyToLayerConf RedisConf

	RWSecProductLock  sync.RWMutex
	secProductInfoMap map[int]*SecProductInfoConf
	idBlackMap        map[int]bool
	ipBlackMap        map[string]bool
	secLimitMgr       *SecLimitMgr
	AccessLimitConf   AccessLimitConf

	WriteLayerToProxyGoroutineNum int
	ReadLayerToProxyGoroutineNum int

	WriteProxyToLayerGoroutineNum int
	ReadProxyToLayerGoroutineNum int

	LogPath string
	LogLevel string

	CookieSecretKey string
}

func NewSecKillConf() *SecKillConf {
	return &SecKillConf{
		secProductInfoMap: make(map[int]*SecProductInfoConf,1024),
	}
}
