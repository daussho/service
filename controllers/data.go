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
    Id int
    Nama string	
    Kapasitas int
    Tersedia int
    Tarif int
}

func init() {
	// register model
	orm.RegisterModel(new(Parkiran))

	// set default database
 	orm.RegisterDataBase("default", "mysql", "root:@/parkir?charset=utf8")

 	// create table
	orm.RunSyncdb("default", false, true)
}

func (this *DataController) GetData() {
	//var hasil Parkiran
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

func getDB(id int){

}