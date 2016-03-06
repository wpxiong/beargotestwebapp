package main

import (
  "github.com/wpxiong/beargo/log"
  "github.com/wpxiong/beargo/webapp"
  "github.com/wpxiong/beargo/appcontext"
  "runtime"
)


func main() {
   log.InitLogWithLevel("Debug")
   runtime.GOMAXPROCS(runtime.NumCPU())
   config := appcontext.AppConfigContext{Port :9001,ConfigPath : "./setting.conf"}
   var appCon appcontext.AppContext = appcontext.AppContext{ ConfigContext :  &config}
   app := webapp.New(&appCon)
   indexCtrl := &IndexControl{}
   sampleCtrl := &SampleControl{}
   app.AddRoute("/test/<pam:[0-9]+>",indexCtrl,"Index",Indexform{})
   app.AddRoute("/sample/<pam:[0-9]+>",sampleCtrl,"Index",Sampleform{})
   app.Start()
}