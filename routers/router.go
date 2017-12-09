package routers

import (
	"service/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/data", &controllers.DataController{})
}
