package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type DataController struct {
	beego.Controller
}

type Parkiran struct {
    Id 			int 	`json:"id"`
    Nama 		string	`json:"nama"`
    Kapasitas 	int 	`json:"kapasitas"`
    Tersedia 	int 	`json:"tersedia"`
    Tarif 		int 	`json:"tarif"`
}

func init() {
	// register model
	orm.RegisterModel(new(Parkiran))	
}

func (this *DataController) GetData() {
	var parkiran Parkiran

	o := orm.NewOrm()
	id := this.Ctx.Input.Param(":id")

	if (id == "1" || id == "2" || id == "3"){
		if (id == "1"){
			o.Raw("SELECT * FROM parkiran WHERE id = ?", 1).QueryRow(&parkiran)
			//parkiran := Parkiran{hasil.Id, hasil.Nama, hasil.Kapasitas, hasil.Tarif}
		} else if (id == "2") {
			o.Raw("SELECT * FROM parkiran WHERE id = ?", 2).QueryRow(&parkiran)
		} else {
			o.Raw("SELECT * FROM parkiran WHERE id = ?", 3).QueryRow(&parkiran)
		}
		//parkiran := Parkiran{1, "Sipil", 500, 2000}
		this.Data["json"] = &parkiran
    	this.Data["tes"] = &parkiran
    	//this.TplName = "blank.html"	
    	this.ServeJSON()
	} else {
		this.TplName = "blank.html"	
	}
	
}