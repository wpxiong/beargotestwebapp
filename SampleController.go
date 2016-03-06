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
  Pam  int8
  Pam2  uint8
  Pam3  float32
  Test  string
  Xiong  []string
}


type SampleControl struct {
  controller.Controller
}

func (*SampleControl) Index(ctx *appcontext.AppContext,form interface{}){
  sampleForm := form.(*Sampleform)
  log.Debug(sampleForm)
  fmt.Fprint(*(ctx.Writer.HttpResponseWriter),"Sample Web")
}
