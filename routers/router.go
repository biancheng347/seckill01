package routers

import (
	"seckill01/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/seckill", &controllers.SecKillController{}, "*:SecKillProduct")
}
