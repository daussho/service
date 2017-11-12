package main

import (
	_ "service/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

