package models

import (
	"fmt"
	"time"
)

type SecResult struct {
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
		return
	}

	data, code, err = secInfoByIf(req.ProductId)
	if err != nil {
		return
	}

	return

}
