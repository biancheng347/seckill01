package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"seckill01/models"
	"seckill01/structModel"
	"strings"
	"time"
)

type SecKillController struct {
	beego.Controller
}

const (
	Code    = "code"
	Data    = "data"
	Message = "message"
)

const (
	CodeDefault    = 200
	DataDefault    = ""
	MessageDefalut = ""
)

func (s *SecKillController) getMulitString(keys ...string) (kv map[string]string, err error) {
	kv = make(map[string]string)
	for _, key := range keys {
		value := s.GetString(key)
		if len(value) == 0 {
			err = fmt.Errorf("parse key : %v ,is failed", key)
			return
		}
		kv[key] = value
	}
	return
}

func (s *SecKillController) getMulitInt(keys ...string) (kv map[string]int, err error) {
	kv = make(map[string]int)
	value := 0
	for _, key := range keys {
		value, err = s.GetInt(key)
		if err != nil {
			return
		}
		kv[key] = value
	}
	return
}

func (s *SecKillController) SecKillProduct() {
	result := make(map[string]interface{})

	result[Code] = CodeDefault
	result[Message] = MessageDefalut
	defer func() {
		s.Data["json"] = result
		s.ServeJSON()
	}()

	mapStrings, err := s.getMulitString("src", "authcode", "time", "nance")
	mapInts, err1 := s.getMulitInt("user_id", "productId")
	if err != nil || err1 != nil {
		result[Code] = 404
		result[Message] = "参数解析错误"
	}

	secRequest := structModel.NewSecRequest()

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

	userAuthSign := s.Ctx.GetCookie("userAuthSign")
	if len(userAuthSign) > 0 {
		secRequest.UserAuthSign = userAuthSign
	}

	secRequest.AccessTime = time.Now()

	addr := s.Ctx.Request.RemoteAddr
	if len(addr) > 0 {
		addrSplit := strings.Split(addr, ":")
		if len(addrSplit) > 0 {
			secRequest.ClientAddr = addrSplit[0]
		}
	}

	fmt.Println("client request ", secRequest)

	data, code, err := models.SecKill(secRequest)
	if err != nil {
		result[Code] = code
		result[Message] = "invalid product Id"
		return
	}

	result[Code] = code
	result[Data] = data
}
