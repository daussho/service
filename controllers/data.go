package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type DataController struct {
	beego.Controller
}

func (c *DataController) Get() {
	c.TplName = "home.html"
}