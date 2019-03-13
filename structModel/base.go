package structModel

import (
	"fmt"
	"github.com/astaxie/beego"
)

var (
	appconfg = beego.AppConfig
)

//other
func appConfigString(key string) (str string, err error) {
	str = appconfg.String(key)
	if len(str) == 0 {
		err = fmt.Errorf("app config string failed,key: %v", key)
		return
	}
	return
}

func appConfigInt(key string) (i int, err error) {
	i, err = appconfg.Int(key)
	if err != nil {
		err = fmt.Errorf("app config int failed,key: %v", key)
		return
	}
	return
}

func appConfigIntValue(num *int, key string) (err error) {
	i, err := appConfigInt(key)
	if err != nil {
		return
	}
	*num = i
	return
}

func appConfigStringValue(str *string, key string) (err error) {
	s, err := appConfigString(key)
	if err != nil {
		return
	}
	*str = s
	return
}




//int
func AppConfigIntValue(num *int, key string) (err error) {
	return appConfigIntValue(num,key)
}

//string
func AppConfigStringValue(str *string, key string) (err error) {
	return appConfigStringValue(str,key)
}