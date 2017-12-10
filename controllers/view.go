package controllers

import (
	"github.com/astaxie/beego"
)

type ViewController struct {
	beego.Controller
}

type Lokasi struct{
	Lok string `form:"loc"`
}



func (c *ViewController) Post() {
	l := Lokasi{}
	c.ParseForm(&l)
	c.Data["lokasi"] = l.Lok

	c.TplName = "view.html"
}

