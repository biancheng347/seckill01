package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"seckill01/models"
	"seckill01/base"
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

	secRequest := base.SecRequstForDic(s.Ctx,mapStrings,mapInts)
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
