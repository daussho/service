package controllers

import (
	"github.com/astaxie/beego"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MainController struct {
	beego.Controller
}

func setDB() *sql.DB {
	db, err:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/kantin")
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	return db
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "home.html"
}