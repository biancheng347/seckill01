package base

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
	ResultChan chan *SecResult
	CloseNotify <-chan bool
}

func NewSecRequest() *SecRequest {
	return &SecRequest{}
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
