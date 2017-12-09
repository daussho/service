package routers

import (
	"service/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/view", &controllers.MainController{})
}
