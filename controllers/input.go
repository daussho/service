package controllers

import(
	"github.com/astaxie/beego"
)

type InputController struct {
	beego.Controller
}

func (this *InputController) Get() {
	this.TplName = "inputdata.html"
}