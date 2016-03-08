package main

import (
  "github.com/wpxiong/beargo/log"
  "github.com/wpxiong/beargo/webapp"
  "github.com/wpxiong/beargo/appcontext"
  "runtime"
  "reflect"
  "fmt"
)

type Config struct {
    Name string
    Meta struct {
        Desc string
        Properties map[string]string
        Users []string
    }
}

func initializeStruct(t reflect.Type, v reflect.Value) {
  for i := 0; i < v.NumField(); i++ {
    f := v.Field(i)
    ft := t.Field(i)
    switch ft.Type.Kind() {
    case reflect.Map:
      f.Set(reflect.MakeMap(ft.Type))
    case reflect.Slice:
      f.Set(reflect.MakeSlice(ft.Type, 0, 0))
    case reflect.Chan:
      f.Set(reflect.MakeChan(ft.Type, 0))
    case reflect.Struct:
      initializeStruct(ft.Type, f)
    case reflect.Ptr:
      fv := reflect.New(ft.Type.Elem())
      initializeStruct(ft.Type.Elem(), fv.Elem())
      f.Set(fv)
    default:
    }
  }
}




func main() {
   log.InitLogWithLevel("Debug")
  
    t := reflect.TypeOf(Config{})
    v := reflect.New(t)
    initializeStruct(t, v.Elem())
    c := v.Interface().(*Config)
    c.Meta.Properties["color"] = "red" // map was already made!
    c.Meta.Users = append(c.Meta.Users, "srid") // so was the slice.
    fmt.Println(v.Interface())
  
   runtime.GOMAXPROCS(runtime.NumCPU())
   config := appcontext.AppConfigContext{Port :9001,ConfigPath : "./setting.conf"}
   var appCon appcontext.AppContext = appcontext.AppContext{ ConfigContext :  &config}
   app := webapp.New(&appCon)
   indexCtrl := &IndexControl{}
   sampleCtrl := &SampleControl{}
   app.AddRoute("/test/<pam:[0-9]+>",indexCtrl,"Index",Indexform{})
   app.AddRoute("/auth/api",indexCtrl,"Index",Indexform{})
   app.AddRoute("/sample/<pam:int>",sampleCtrl,"Index",Sampleform{})
   app.AddRoute("/sample/<pam:int>/create",sampleCtrl,"Create",Sampleform{})
   app.Start()
}