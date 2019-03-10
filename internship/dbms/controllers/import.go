package controllers

import (
	"dbms/models"
	"github.com/astaxie/beego"
)

type ImportController struct {
	beego.Controller
}

func (this *ImportController) Get() {
    this.Data["Import"] = "beego.me"
	this.TplName = "import.html"

	managements,err := models.GetAllIncrease()
	if err != nil{
		beego.Error(err)
	} else{
		this.Data["Managements"] = managements
	}

	imports,err := models.GetAllImport()
	if err != nil{
		beego.Error(err)
	} else{
		this.Data["Imports"] = imports
	}
}

func(this *ImportController) Post(){
	types := this.Input().Get("types")
	names := this.Input().Get("names")
	modelss := this.Input().Get("models")
	manufacturers := this.Input().Get("manufacturers")
	costs := this.Input().Get("costs")
	importnumbers := this.Input().Get("importnumbers")
	tid := this.Input().Get("id")
	stock := this.Input().Get("export")

	err1 := models.ImportCommodity(types,names,modelss,manufacturers,costs,importnumbers)
	err2 := models.IncreaseStock(tid,importnumbers,stock)

	if err1 != nil{
		beego.Error(err1)
	}

	if err2 != nil{
		beego.Error(err2)
	}
	this.Redirect("/import",301)
}

func(this *ImportController) Modify() {
	this.TplName="increases_modify.html"

	tid:=this.Input().Get("tid")
	topic,err :=models.GetManagement(tid)
	if err !=nil{
		beego.Error(err)
		this.Redirect("/import",302)
		return
	}

	this.Data["Management"] = topic
	this.Data["Tid"] = tid
} 