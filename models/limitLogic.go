package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"seckill01/structModel"
	"time"
)

func antiSpam(req *structModel.SecRequest) (err error) {

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

	if seckillconf.SecProductInfoMap == nil {
		code = ErrClientClosed
		err = fmt.Errorf("sec product info map is nil")
		return
	}
	v, ok := seckillconf.SecProductInfoMap[productId]
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

func SecKill(req *structModel.SecRequest) (data map[string]interface{}, code int, err error) {
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

	if code != 0 {
		logs.Warn("userId: %d secInfByid failed, code[%d] req[%v]", req.UserId, code, req)
		return
	}

	userKey := fmt.Sprintf("%s_%s", req.UserId, req.ProductId)
	seckillconf.UserConnMap[userKey] = req.ResultChan
	seckillconf.SecReqChan <- req
	defer func() {
		seckillconf.UserConnMapLock.Lock()
		delete(seckillconf.UserConnMap, userKey)
		seckillconf.UserConnMapLock.Unlock()
	}()

	data ,code ,err = req.ReqSelect()
	return
}
