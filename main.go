package main

import (
  "github.com/wpxiong/beargo/log"
  "github.com/wpxiong/beargo/webapp"
  "github.com/wpxiong/beargo/appcontext"
  "github.com/wpxiong/beargo/moudle"
  "time"
  "runtime"
)



type AddressInformation struct {
  Addressid int `id:"true"`
  Userid int
  AddressName string  `notnull:"true"     length:"128" `
}

type ClassInfo struct {
  ClassName string `length:"128"`
  ClassId int `id:"true"   auto_increment:"true"`
  Address  string `length:"128"  notnull:"true"`
  Goup    []GroupInfo `relation_type:"onetomany"  column_name:"classid" referenced_column_name:"classid"`
}


type GroupInfo struct {
  UserId  int  `id:"true"`
  ClassId int  `id:"true"`
  Class ClassInfo  `relation_type:"manytoone" column_name:"classid"  referenced_column_name:"classid"`
  User  UserInfo  `relation_type:"manytoone"   column_name:"userid" referenced_column_name:"id"`
}


type UserInfo struct {
  Id      int `id:"true"`
  UserName string   `id:"true"     notnull:"true"     length:"128"     default_value:"testping"`
  UserAge  uint16  `default_value:"54"`
  UserSex  bool  `default_value:"true"`
  Test     complex64 `default_value:"12.34,11.78"`
  CreateTime  time.Time
  Addressid  int
  Goup    []GroupInfo `relation_type:"onetomany" column_name:"id" referenced_column_name:"userid"`
  Address  AddressInformation `relation_type:"onetoone" column_name:"addressid" referenced_column_name:"addressid"`
}


func main() {
   log.InitLogWithLevel("Debug")
   configMap := InitConfig()
   
   moudleInstance :=  moudle.CreateModuleInstance(moudle.MYSQL,"test","tcp(localhost:3306)","","") 
   
   moudleInstance.AddTable(UserInfo{})
   moudleInstance.AddTable(ClassInfo{})
   moudleInstance.AddTable(GroupInfo{})
   moudleInstance.AddTable(AddressInformation{})

   moudleInstance.InitialDB(true)
   moudleInstance.Insert(AddressInformation{Addressid:1,Userid:12,AddressName:"tokyo shibuya"})
   moudleInstance.Insert(UserInfo{Addressid:1,UserName:"xiong",UserAge:22,UserSex:true,Test:complex(23.4,56.7),CreateTime:time.Now()})
   
   var info moudle.QueryInfo  = moudleInstance.Query(UserInfo{},moudle.EAGER,[]string{"Goup","Address"})
   info.GetResultList()
   runtime.GOMAXPROCS(runtime.NumCPU())
   config := appcontext.AppConfigContext{Port :9001,ConfigPath : "./setting.conf"}
   
   var appCon appcontext.AppContext = appcontext.AppContext{ ConfigContext :  &config}
   
   app := webapp.New(&appCon,configMap)
   
   indexCtrl := &IndexControl{}
   sampleCtrl := &SampleControl{}
   app.AddAutoRoute("/<name:.*>/<password:[0-9]+>",indexCtrl,Indexform{})
   app.AddAutoRouteWithViewPath("/sample/<pam:int>/",sampleCtrl,Sampleform{},"/sampletest")
   
   app.Start()
}