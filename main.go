package main

import (
	"github.com/astaxie/beego"
	"seckill01/config"
	_ "seckill01/routers"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		panic(err)
		return
	}
	beego.Run()
}
