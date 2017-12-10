package controllers

import(
	"time"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type TiketController struct {
	beego.Controller
}

type Tiket struct {
    Id 				int
    TiketId 		int
    TanggalMasuk 	time.Time
    TanggalKeluar 	time.Time
}


func init() {
	// register model
	orm.RegisterModel(new(Tiket))	
}

func (this *TiketController) GetData() {

}