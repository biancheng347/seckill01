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

type SecKillConf struct {
	RedisBlackConf RedisConf

	RWSecProductLock  sync.RWMutex
	secProductInfoMap map[int]*SecProductInfoConf
}

func NewSecKillConf() *SecKillConf {
	return &SecKillConf{}
}
