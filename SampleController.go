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

type  Sampleform struct{
  Name string
}


type SampleControl struct {
  controller.Controller
}

func (*SampleControl) Index(ctx *appcontext.AppContext,form interface{}){
  
  
  fmt.Fprint(*(ctx.Writer.HttpResponseWriter),"Sample Web")
}
