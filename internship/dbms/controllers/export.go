package controllers

import (
	"dbms/models"
	"github.com/astaxie/beego"
//	"github.com/tealeg/xlsx"
)

type ExportController struct {
	beego.Controller
}

func (this *ExportController) Get() {
	this.Data["Export"] = "beego.me"
	this.TplName = "export.html"

/*	excelFileName := "/Users/sceaj/Desktop/test.xlsx"
    xlFile, error := xlsx.OpenFile(excelFileName)
    if error != nil {
        beego.Error(error)
        return
    }
    for _, sheet := range xlFile.Sheets {
        for _, row := range sheet.Rows {
            for _, cell := range row.Cells {
                text := cell.String()
            }
        }
    } */

	managements,err := models.GetAllIncrease()
	if err != nil{
		beego.Error(err)
	} else{
		this.Data["Managements"] = managements
	}

	exports,err := models.GetAllExport()
	if err != nil{
		beego.Error(err)
	} else{
		this.Data["Exports"] = exports
	}
}

func(this *ExportController) Post(){
	tid := this.Input().Get("id")
	types := this.Input().Get("types")
	names := this.Input().Get("names")
	modelss := this.Input().Get("modelss")
	prices := this.Input().Get("prices")
	exports := this.Input().Get("exports")
	stock := this.Input().Get("export")

	err1 := models.ExportCommodity(types,names,modelss,prices,exports)
	err2 := models.DeleteStock(tid,exports,stock)

	if err1 != nil{
		beego.Error(err1)
	}

	if err2 != nil{
		beego.Error(err2)
	}
	this.Redirect("/export",301)
}


func(this *ExportController) Modify() {
	this.TplName="export_modify.html"

	tid:=this.Input().Get("tid")
	topic,err :=models.GetManagement(tid)
	if err !=nil{
		beego.Error(err)
		this.Redirect("/export",302)
		return
	}

	this.Data["Management"] = topic
	this.Data["Tid"] = tid
} 