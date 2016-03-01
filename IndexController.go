package main

import (
  "github.com/wpxiong/beargo/log"
  "github.com/wpxiong/beargo/controller"
  "github.com/wpxiong/beargo/route"
  "fmt"
  "net/http"
)

func init() {
  log.InitLog()
}

type IndexControl struct {
  controller.Controller
}

func (*IndexControl) Index(rti *route.RouteInfo){
  var w http.ResponseWriter = *rti.Writer.HttpResponseWriter
  fmt.Fprintf(w,"Edit")
}
