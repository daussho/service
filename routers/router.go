package routers

import (
	"service/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/data/:id([1-3]+", &controllers.DataController{})
    beego.Router("/view", &controllers.ViewController{})
}
