package controllers

import (
	"dbms/models"
	"github.com/astaxie/beego"
)

type WarehouseController struct {
	beego.Controller
}

func (this *WarehouseController) Get() {
	this.Data["WareHouse"] = "beego.me"
	this.TplName = "warehouse.html"

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