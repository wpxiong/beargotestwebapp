package main

import (
  "github.com/wpxiong/beargo/log"
  "github.com/wpxiong/beargo/webapp"
  "github.com/wpxiong/beargo/appcontext"
  "runtime"
)



func main() {
   log.InitLogWithLevel("Debug")
   configMap := InitConfig()
   
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