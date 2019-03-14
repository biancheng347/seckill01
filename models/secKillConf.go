package models

import (
	"github.com/garyburd/redigo/redis"
	"seckill01/structModel"
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
	RedisBlackConf        structModel.RedisConf
	BlackRedisPool        *redis.Pool
	idBlackMap        map[int]bool //
	ipBlackMap        map[string]bool //
}

type LayerToProxyConf struct {
	RedisLayerToProxyConf structModel.RedisConf
	LayerToProxyRedisPool *redis.Pool
	WriteLayerToProxyGoroutineNum int
	ReadLayerToProxyGoroutineNum  int
}

//type ProxyToLayerConf struct{
//	RedisProxyToLayerConf structModel.RedisConf
//	ProxyToLayerRedisPool *redis.Pool
//	WriteProxyToLayerGoroutineNum int
//	ReadProxyToLayerGoroutineNum  int
//}


type SecKillConf struct {
	BlackConf
	LayerToProxyConf
	RedisProxyToLayerConf structModel.RedisConf
	ProxyToLayerRedisPool *redis.Pool
	WriteProxyToLayerGoroutineNum int
	ReadProxyToLayerGoroutineNum  int
	
	EtcdConf structModel.EtcdConf





	Logs structModel.LogsConf

	RWSecProductLock  sync.RWMutex
	SecProductInfoMap map[int]*SecProductInfoConf //

	SecReqChan     chan *structModel.SecRequest//
	SecReqChanSize int

	UserConnMap     map[string]chan *structModel.SecResult //
	UserConnMapLock sync.Mutex

	CookieSecretKey string
	ReferWhiteList []string
	AccessLimitConf  structModel.AccessLimitConf
	secLimitMgr       *SecLimitMgr //

}

func NewSecKillConf() *SecKillConf {
	return &SecKillConf{
		SecProductInfoMap: make(map[int]*SecProductInfoConf, 1024),
		//idBlackMap: make(map[int]bool,10000),
		//ipBlackMap: make(map[string]bool,10000),
		secLimitMgr: &SecLimitMgr{
			UserLimitMap:make(map[int]*Limit,10000),
			IpLimitMap:make(map[string]*Limit,10000),
		},
		SecReqChan: make(chan *structModel.SecRequest,10000),
		UserConnMap: make(map[string]chan *structModel.SecResult,10000),
	}
}

