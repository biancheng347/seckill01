package models

import (
	"github.com/garyburd/redigo/redis"
	"seckill01/structModel"
	"sync"
	"time"
)

type SecResult struct {
	ProductId int
	UserId    int
	Code      int
	Token     string
}

type SecRequest struct {
	ProductId    int
	Source       string
	AuthCode     string
	SecTime      string
	Nance        string
	UserId       int
	UserAuthSign string
	AccessTime   time.Time
	ClientAddr   string
}

func NewSecRequest() *SecRequest {
	return &SecRequest{}
}

//type RedisConf struct {
//	RedisAddr        string
//	RedisMaxIdle     int
//	RedisMaxActive   int
//	RedisIdleTimeout int
//}

type SecProductInfoConf struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int
	Total     int
	Remain    int
}

//type AccessLimitConf struct {
//	IPSecAccessLimit   int
//	UserSecAccessLimit int
//	IPMinAccessLimit   int
//	UserMinAccessLimit int
//}

//type EtcdConfParam struct{
//	EtcdAddr          string
//	Timeout           int
//}
//
//type EtcdConf struct{
//	EtcdConfParam
//	EtcdSecKeyPrefix  string
//	EtcdSecProductKey string
//}


type SecKillConf struct {
	RedisBlackConf        structModel.RedisConf
	RedisLayerToProxyConf structModel.RedisConf
	RedisProxyToLayerConf structModel.RedisConf

	EtcdConf structModel.EtcdConf

	BlackRedisPool        *redis.Pool
	ProxyToLayerRedisPool *redis.Pool
	LayerToProxyRedisPool *redis.Pool

	WriteLayerToProxyGoroutineNum int
	ReadLayerToProxyGoroutineNum  int
	WriteProxyToLayerGoroutineNum int
	ReadProxyToLayerGoroutineNum  int

	idBlackMap        map[int]bool //
	ipBlackMap        map[string]bool //

	Logs structModel.LogsConf

	RWSecProductLock  sync.RWMutex
	SecProductInfoMap map[int]*SecProductInfoConf //

	SecReqChan     chan *SecRequest//
	SecReqChanSize int

	UserConnMap     map[string]chan *SecResult //
	UserConnMapLock sync.Mutex

	CookieSecretKey string
	ReferWhiteList []string
	AccessLimitConf  structModel.AccessLimitConf
	secLimitMgr       *SecLimitMgr //

}

func NewSecKillConf() *SecKillConf {
	return &SecKillConf{
		SecProductInfoMap: make(map[int]*SecProductInfoConf, 1024),
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

