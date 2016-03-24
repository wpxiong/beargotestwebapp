package main

import (
  "github.com/wpxiong/beargo/log"
  "github.com/wpxiong/beargo/webapp"
  "github.com/wpxiong/beargo/appcontext"
  "github.com/wpxiong/beargo/moudle"
  "time"
  "runtime"
)

type AddInfo struct {
  Address string  `length:"12"`
  Id      int `id:"true"`
}

type UserInfo struct {
  Id      int `id:"true"`
  UserName string   `id:"true"     notnull:"true"     length:"128"     default_value:"testping"`
  UserAge  uint16  `default_value:"54"`
  UserSex  bool  `default_value:"true"`
  Test     complex64 `default_value:"12.34,11.78"`
  Add      []AddInfo
  CreateTime  time.Time 
}


func main() {
   log.InitLogWithLevel("Debug")
   configMap := InitConfig()
   
   moudleInstance :=  moudle.CreateModuleInstance(moudle.MYSQL,"test","tcp(localhost:3306)","","")
   
   moudleInstance.AddTable(UserInfo{})
   moudleInstance.AddTable(AddInfo{})
   
   moudleInstance.InitialDB(true)
   
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