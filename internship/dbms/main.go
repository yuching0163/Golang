package main

import (
	_ "dbms/routers"
	"dbms/controllers"
	"dbms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init(){
	//註冊數據庫
	models.RegisterDB()
}

func main() {
    //開啟ORM調適模式
	orm.Debug = true
	//自動建表
	orm.RunSyncdb("default",false,true)

    //註冊beego路由
	beego.Router("/",&controllers.MainController{})
    beego.Router("/import",&controllers.ImportController{})
    beego.Router("/export",&controllers.ExportController{})
    beego.AutoRouter(&controllers.ImportController{})
    beego.AutoRouter(&controllers.ExportController{})
    beego.Router("/increases",&controllers.IncreaseController{})
    beego.Router("/warehouse",&controllers.WarehouseController{})
    beego.Router("/login",&controllers.LoginController{})



	//啟動beego
	beego.Run()
}

