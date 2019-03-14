package base

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"strings"
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
	ResultChan chan *SecResult
	CloseNotify <-chan bool
}

func NewSecRequest() *SecRequest {
	return &SecRequest{}
}

func SecReeustForDic(ctx *context.Context,mapStrings map[string]string,mapInts map[string]int) (secRequest *SecRequest) {
	secRequest = NewSecRequest()
	if source, ok := mapStrings["src"]; ok {
		secRequest.Source = source
	}

	if authcode, ok := mapStrings["authcode"]; ok {
		secRequest.AuthCode = authcode
	}

	if secTime, ok := mapStrings["time"]; ok {
		secRequest.SecTime = secTime
	}

	if nance, ok := mapStrings["nance"]; ok {
		secRequest.Nance = nance
	}

	if productId, ok := mapInts["productId"]; ok {
		secRequest.ProductId = productId
	}

	if userId, ok := mapInts["user_id"]; ok {
		secRequest.UserId = userId
	}

	userAuthSign := ctx.GetCookie("userAuthSign")
	if len(userAuthSign) > 0 {
		secRequest.UserAuthSign = userAuthSign
	}

	secRequest.AccessTime = time.Now()

	addr := ctx.Request.RemoteAddr
	if len(addr) > 0 {
		addrSplit := strings.Split(addr, ":")
		if len(addrSplit) > 0 {
			secRequest.ClientAddr = addrSplit[0]
		}
	}
	return
}


func (p *SecRequest)ReqSelect() (data map[string]interface{}, code int, err error) {
	if data == nil {
		return
	}
	ticker := time.NewTicker(time.Second * 10)
	defer func() {
		ticker.Stop()
	}()

	select {
	case <-ticker.C:
		//code = models.ErrProcessTimeout
		code = 1008
		err = fmt.Errorf("request timtout")
	case <-p.CloseNotify:
		//code = models.ErrClientClosed
		code = 1009
		err = fmt.Errorf("client alread close")
	case result := <-p.ResultChan:
		code = result.Code
		data["productId"] = result.ProductId
		data["token"] = result.Token
		data["user_id"] = result.UserId
	}
	return
}
