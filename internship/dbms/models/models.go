package models

import(
  "time"
  "os"
  "path"
  "strconv"
  "github.com/Unknwon/com"
  "github.com/astaxie/beego/orm"
  _"github.com/go-sql-driver/mysql"
  )

const(
  _DB_NAME ="root:1234@tcp(127.0.0.1:3306)/dbms?charset=utf8&loc=Asia%2FShanghai"
  _MYSQL_DRIVE="mysql"
)

type Management struct {
  Id            int64
  Type          string
  Name          string
  Model         string
  Manufacturer  string
  Cost          int64
  Import        int64
  Price         int64
  Export        int64
  Stock         int64
  ImportCreated time.Time `orm:"index"`
  ExportCreated time.Time `orm:"index"`
} 

type Import struct {
  Id            int64
  Type          string
  Name          string
  Model         string
  Manufacturer  string
  Cost          int64
  Import        int64
  ImportCreated time.Time `orm:"index"`
} 

type Export struct {
  Id            int64
  Type          string
  Name          string
  Model         string
  Price         int64
  Export        int64
  ExportCreated time.Time `orm:"index"`
} 

func RegisterDB(){
  if !com.IsExist(_DB_NAME){
    os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
    os.Create(_DB_NAME)
  }

  orm.RegisterModel(new(Management))
  orm.RegisterModel(new(Import))
  orm.RegisterModel(new(Export))
  orm.RegisterDriver("mysql", orm.DRMySQL)
  orm.RegisterDataBase("default",_MYSQL_DRIVE,_DB_NAME,10)
}

func IncreaseCommodity(original,name,model,cost,exportnumber string) error {
  o := orm.NewOrm()
    costNum,cerr := strconv.ParseInt(cost,10,64)
    exportnumberNum,eerr := strconv.ParseInt(exportnumber,10,64)
    if cerr != nil{
      return cerr
    }
    if eerr != nil{
      return eerr
    }
    management := &Management{
      Type: original,
      Name: name,
      Model: model,
      Cost: costNum,
      Price: exportnumberNum,
      ImportCreated: time.Now(),
  }
  _,err := o.Insert(management)
  return err
}

func ImportCommodity(types,names,models,manufacturers,costs,importnumbers string) error {
  o := orm.NewOrm()
    costsNum,cerr := strconv.ParseInt(costs,10,64)
    importnumbersNum,cerr := strconv.ParseInt(importnumbers,10,64)
    if cerr != nil{
      return cerr
    }
    imports := &Import{
      Type: types,
      Name: names,
      Model: models,
      Manufacturer: manufacturers,
      Cost: costsNum,
      Import: importnumbersNum,
      ImportCreated: time.Now(),
  }
  _,err1 := o.Insert(imports)
  return err1
}

func ExportCommodity(types,names,modelss,prices,exports string) error {
  o := orm.NewOrm()
    pricesNum,cerr := strconv.ParseInt(prices,10,64)
    exportsNum,cerr := strconv.ParseInt(exports,10,64)
    if cerr != nil{
      return cerr
    }
    export := &Export{
      Type: types,
      Name: names,
      Model: modelss,
      Price: pricesNum,
      Export: exportsNum,
      ExportCreated: time.Now(),
  }
  _,err1 := o.Insert(export)
  return err1
}

  func IncreaseStock(tid,importnumbers,stock string) error{
  tidNum,err:=strconv.ParseInt(tid,10,64)
  stockNum,err1:=strconv.ParseInt(stock,10,64)
  importnumberNum,err2:=strconv.ParseInt(importnumbers,10,64)

  if err!=nil{
    return err
  }
  if err1!=nil{
    return err
  }
  if err2!=nil{
    return err
  }
  o := orm.NewOrm()

  sum := stockNum + importnumberNum
  management := &Management{Id:tidNum}
  if o.Read(management) == nil{
    management.Stock = sum
    o.Update(management)
  }
  return err
}

  func DeleteStock(tid,exports,stock string) error{
  tidNum,err:=strconv.ParseInt(tid,10,64)
  stockNum,err1:=strconv.ParseInt(stock,10,64)
  exportsNum,err2:=strconv.ParseInt(exports,10,64)

  if err!=nil{
    return err
  }
  if err1!=nil{
    return err
  }
  if err2!=nil{
    return err
  }
  o := orm.NewOrm()

  sum := stockNum - exportsNum
  management := &Management{Id:tidNum}
  if o.Read(management) == nil{
    management.Stock = sum
    o.Update(management)
  }
  return err
}


func GetAllIncrease()([]*Management,error){
  o := orm.NewOrm()

  managements := make([]*Management,0)

  qs := o.QueryTable("management")
  _,err := qs.All(&managements)
  return managements,err
}

func GetAllImport()([]*Import,error){
  o := orm.NewOrm()

  imports := make([]*Import,0)

  qs := o.QueryTable("import")
  _,err := qs.All(&imports)
  return imports,err
}

func GetAllExport()([]*Export,error){
  o := orm.NewOrm()

  exports := make([]*Export,0)

  qs := o.QueryTable("export")
  _,err := qs.All(&exports)
  return exports,err
}

func DelCommodity(id string) error {
  cid,err := strconv.ParseInt(id,10,64)
  if err != nil{
    return err
  }

  o:=orm.NewOrm()

  commodity := &Management{Id:cid}
  _,err = o.Delete(commodity)
  return err
}

func GetManagement(tid string) (*Management,error){
  tidNum,err:=strconv.ParseInt(tid,10,64)
  if err!=nil{
    return nil,err
  }

  o := orm.NewOrm()

  management := new(Management)

  qs := o.QueryTable("management")
  err = qs.Filter("id",tidNum).One(management)
  if err !=nil{
    return nil,err
  }   
  return management,err
}