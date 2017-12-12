package controllers

import(
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type InputController struct {
	beego.Controller
}

type Tiket struct {
    Id 				int
    TiketId 		int
    TanggalMasuk 	time.Time
    TanggalKeluar 	time.Time
}

type Form struct{
	Tipe string `form:"tipe"`
	Lok string `form:"loc"`
	Id string `form:"id"`
}

func init() {
	// register model
	orm.RegisterModel(new(Tiket))	
}

func (c *InputController) Get() {
	c.TplName = "inputdata.html"
}

func (c *InputController) Post() {
	var sisa int

	o := orm.NewOrm()

	l := Form{}
	c.ParseForm(&l)

	if (l.Tipe == "input"){
		o.Raw("SELECT tersedia FROM parkiran WHERE id = ?", l.Lok).QueryRow(&sisa)
		o.Raw("INSERT INTO tiket (id, tiket_id, tanggal_masuk, tanggal_keluar) VALUES (?, NULL, CURRENT_TIMESTAMP, NULL)", l.Lok).Exec()
		sisa--;
		o.Raw("UPDATE parkiran SET tersedia = ? WHERE id = ?", sisa, l.Lok).Exec()	
	} else if (l.Tipe == "update"){
		o.Raw("SELECT tersedia FROM parkiran WHERE id = ?", l.Lok).QueryRow(&sisa)
		o.Raw("UPDATE tiket SET tanggal_keluar = CURRENT_TIMESTAMP WHERE tiket_id = ?", l.Id).Exec()
		sisa++;
		o.Raw("UPDATE parkiran SET tersedia = ? WHERE id = ?", sisa, l.Lok).Exec()	
	}
	
	

	c.TplName = "inputdata.html"
}