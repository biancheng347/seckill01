package models

import (
	"github.com/garyburd/redigo/redis"
	"seckill01/base"
	"sync"
)

//type SecResult struct {
//	ProductId int
//	UserId    int
//	Code      int
//	Token     string
//}
//
//type SecRequest struct {
//	ProductId    int
//	Source       string
//	AuthCode     string
//	SecTime      string
//	Nance        string
//	UserId       int
//	UserAuthSign string
//	AccessTime   time.Time
//	ClientAddr   string
//	ResultChan chan *SecResult
//	CloseNotify <-chan bool
//}

//func NewSecRequest() *SecRequest {
//	return &SecRequest{}
//}

type SecProductInfoConf struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int
	Total     int
	Remain    int
}

type BlackConf struct{
	RedisBlackConf base.RedisConf
	BlackRedisPool *redis.Pool
	idBlackMap     map[int]bool //
	ipBlackMap     map[string]bool //
}

type LayerToProxyConf struct {
	RedisLayerToProxyConf         base.RedisConf
	LayerToProxyRedisPool         *redis.Pool
	WriteLayerToProxyGoroutineNum int
	ReadLayerToProxyGoroutineNum  int
}

type ProxyToLayerConf struct{
	RedisProxyToLayerConf         base.RedisConf
	ProxyToLayerRedisPool         *redis.Pool
	WriteProxyToLayerGoroutineNum int
	ReadProxyToLayerGoroutineNum  int
}

type SecRequet struct  {

}

type UseConn struct {
	UserConnMap     map[string]chan *base.SecResult //
	UserConnMapLock sync.Mutex
}

type SecProduct struct {
	RWSecProductLock  sync.RWMutex
	SecProductInfoMap map[int]*SecProductInfoConf //
}

type SecReqChanConf struct {
	SecReqChan     chan *base.SecRequest //
	SecReqChanSize int
}

type SecLimitConf struct {
	AccessLimitConf base.AccessLimitConf
	secLimitMgr     *SecLimitMgr //
}

type SecKillConf struct {
	BlackConf
	LayerToProxyConf
	ProxyToLayerConf
	
	base.LogsConf

	SecProduct

	SecReqChanConf

	UseConn

	SecLimitConf
	CookieSecretKey string
	ReferWhiteList  []string

	base.EtcdConf
}

func NewSecKillConf() *SecKillConf {
	return &SecKillConf{
		//SecProductInfoMap: make(map[int]*SecProductInfoConf, 1024),
		//idBlackMap: make(map[int]bool,10000),
		//ipBlackMap: make(map[string]bool,10000),
		//secLimitMgr: &SecLimitMgr{
		//	UserLimitMap:make(map[int]*Limit,10000),
		//	IpLimitMap:make(map[string]*Limit,10000),
		//},
		//SecReqChan: make(chan *base.SecRequest,10000),
		//UserConnMap: make(map[string]chan *base.SecResult,10000),
	}
}

