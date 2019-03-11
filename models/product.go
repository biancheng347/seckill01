package models

import "time"

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
