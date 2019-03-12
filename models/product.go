package models

import (
	"fmt"
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

func antiSpam(req *SecRequest) (err error) {

	if seckillconf.idBlackMap == nil || seckillconf.ipBlackMap == nil {
		err = fmt.Errorf("client is closed")
		return
	}

	if _, ok := seckillconf.idBlackMap[req.UserId]; ok {
		err = fmt.Errorf("invalid request")
		fmt.Println("userId is added to idBlackList")
		return
	}

	if _, ok := seckillconf.ipBlackMap[req.ClientAddr]; ok {
		err = fmt.Errorf("invalid request")
		fmt.Println("ip is added to ipBlackList ")
		return
	}

	//uid rate controller

	seckillconf.secLimitMgr.Lock.Lock()
	userLimitMap := seckillconf.secLimitMgr.UserLimitMap
	if userLimitMap == nil {
		err = fmt.Errorf("client is closed")
		return
	}
	limit, ok := userLimitMap[req.UserId]
	if !ok {
		limit = &Limit{
			secLimit: &SecLimit{},
			minLimit: &MinLimit{},
		}
		userLimitMap[req.UserId] = limit
	}

	secIdCount := limit.secLimit.Count(req.AccessTime.Unix())
	midIdCount := limit.minLimit.Count(req.AccessTime.Unix())

	//ip rate
	ipLimitMap := seckillconf.secLimitMgr.IpLimitMap
	if ipLimitMap == nil {
		err = fmt.Errorf("client is closed")
		return
	}
	limit, ok = ipLimitMap[req.ClientAddr]
	if !ok {
		limit = &Limit{
			secLimit: &SecLimit{},
			minLimit: &MinLimit{},
		}
		ipLimitMap[req.ClientAddr] = limit
	}
	secIpCount := limit.secLimit.Count(req.AccessTime.Unix())
	midIpCount := limit.minLimit.Count(req.AccessTime.Unix())

	seckillconf.secLimitMgr.Lock.Unlock()

	if secIdCount > seckillconf.AccessLimitConf.UserSecAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}

	if midIdCount > seckillconf.AccessLimitConf.UserMinAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}

	if secIpCount > seckillconf.AccessLimitConf.IPSecAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}

	if midIpCount > seckillconf.AccessLimitConf.IPMinAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}

	return
}

func secInfoByIf(productId int) (data map[string]interface{}, code int, err error) {
	data = make(map[string]interface{})
	start := false
	end := false
	status := ""

	if seckillconf.secProductInfoMap == nil {
		code = ErrClientClosed
		err = fmt.Errorf("sec product info map is nil")
		return
	}
	v, ok := seckillconf.secProductInfoMap[productId]
	if !ok {
		code = ErrNotFoundProductId
		err = fmt.Errorf("not found this product for id")
		return
	}
	now := time.Now().Unix()
	if now-v.StartTime < 0 {
		code = ErrActiveNotStart
		status = "sec kill is not start"
	} else if now-v.EndTime > 0 {
		start = true
		end = true
		code = ErrActiveAlreadyEnd
		status = "sec kill is end"
	} else if v.Status == ErrActiveSaleOut {
		start = true
		end = true
		code = ErrActiveSaleOut
		status = "sec kill have saled"
	} else if now-v.StartTime > 0 && now-v.EndTime < 0 {
		end = true
		code = SuccActvieDoing
		status = "sec kill is saling"
	}

	data["productId"] = productId
	data["start"] = start
	data["end"] = end
	data["status"] = status
	return
}

func (req *SecRequest) SecKill() (data map[string]interface{}, code int, err error) {
	seckillconf.RWSecProductLock.RLock()
	defer seckillconf.RWSecProductLock.RUnlock()

	if err = antiSpam(req); err != nil {
		code = ErrInvalidRequest
		return
	}

	data, code, err = secInfoByIf(req.ProductId)
	if err != nil {
		return
	}

	return

}
