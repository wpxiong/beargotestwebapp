package main

import (
  "github.com/wpxiong/beargo/log"
  "github.com/wpxiong/beargo/controller"
  "github.com/wpxiong/beargo/appcontext"
  "fmt"
)

func init() {
  log.InitLog()
}

type IndexControl struct {
  controller.Controller
}

func (*IndexControl) Index(ctx *appcontext.AppContext){
  fmt.Fprint(*(ctx.Writer.HttpResponseWriter),"Test App")
}
