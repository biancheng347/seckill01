package structModel

import "time"

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