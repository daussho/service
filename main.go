package main

import (
	_ "service/routers"
	_ "service/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
 	orm.RegisterDataBase("default", "mysql", "root:@localhost/parkir?charset=utf8")
 	orm.RegisterModel(new(Parkiran))
}

func main() {
	beego.Run()
}

