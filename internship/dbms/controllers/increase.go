package controllers

import (
	"dbms/models"
	"github.com/astaxie/beego"
)

type IncreaseController struct {
	beego.Controller
}

func (this *IncreaseController) Get() {
	op:=this.Input().Get("op")

    switch op{
    case "add":
    	break
    case "del":
    	id:=this.Input().Get("id")
    	if len(id) == 0{
    		break
    	}

    	err := models.DelCommodity(id)
    	if err!=nil{
    		beego.Error(err)
    	}
    	this.Redirect("/increases",302)
    	return
    }

	
	this.Data["Increase"] = "beego.me"
	this.TplName = "increases.html"
	managements,err := models.GetAllIncrease()
	if err != nil{
		beego.Error(err)
	} else{
		this.Data["Managements"] = managements
	}
}

func(this *IncreaseController) Post(){
	original := this.Input().Get("original")
	name := this.Input().Get("name")
	model := this.Input().Get("model")
	cost := this.Input().Get("cost")
	exportnumber := this.Input().Get("exportnumber")

	var err error
	err = models.IncreaseCommodity(original,name,model,cost,exportnumber)

	if err != nil{
		beego.Error(err)
	}
	this.Redirect("/increases",302)
}