package main

import (
	_ "service/routers"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

func init() {
	// register model
	//orm.RegisterModel(new(Parkiran))

	// set default database
 	orm.RegisterDataBase("default", "mysql", "root:@/parkir?charset=utf8")

 	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}

