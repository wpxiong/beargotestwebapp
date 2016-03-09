package main

import (
  "github.com/wpxiong/beargo/log"
  "github.com/wpxiong/beargo/webapp"
  "github.com/wpxiong/beargo/appcontext"
  "runtime"
)



func main() {
   log.InitLogWithLevel("Debug")
   funcmap := InitConfig()
   runtime.GOMAXPROCS(runtime.NumCPU())
   config := appcontext.AppConfigContext{Port :9001,ConfigPath : "./setting.conf"}
   
   var appCon appcontext.AppContext = appcontext.AppContext{ ConfigContext :  &config}
   
   app := webapp.New(&appCon,funcmap)
   
   indexCtrl := &IndexControl{}
   sampleCtrl := &SampleControl{}
   app.AddRoute("/<name:.*>/<password:[0-9]+>",indexCtrl,"Index",Indexform{})
   app.AddRoute("/sample/<pam:int>",sampleCtrl,"Index",Sampleform{})
   app.AddRoute("/sample/<pam:int>/create",sampleCtrl,"Create",Sampleform{})
   
   app.Start()
}