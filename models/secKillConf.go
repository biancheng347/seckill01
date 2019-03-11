package models

import "sync"

var (
	seckillconf = NewSecKillConf()
)

type RedisConf struct {
	RedisAddr      string
	RedisMaxIdle   int
	RedisMaxActive int
	RedisIdleTime  int
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
	RedisBlackConf RedisConf

	RWSecProductLock  sync.RWMutex
	secProductInfoMap map[int]*SecProductInfoConf
	idBlackMap        map[int]bool
	ipBlackMap        map[string]bool
	secLimitMgr       *SecLimitMgr
	AccessLimitConf   AccessLimitConf
}

func NewSecKillConf() *SecKillConf {
	return &SecKillConf{}
}
