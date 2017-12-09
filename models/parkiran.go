package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
 	_ "github.com/go-sql-driver/mysql"
)

// User is a wrapper for result row in this example
type Parkiran struct {
    id int
    nama string
    kapasitas int
    //tersedia int
    tarif int
}

func init() {
    orm.RegisterModel(new(Parkiran))
}

func (u *Parkiran) TempatParkir() Parkiran {
	o := orm.NewOrm()
	parkiran := models.Parkiran{Id: 1}
	err := o.Read(&parkiran)
	return parkiran
}