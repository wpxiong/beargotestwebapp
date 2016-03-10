package main

import (
  "github.com/wpxiong/beargo/log"
  "github.com/wpxiong/beargo/controller"
  "github.com/wpxiong/beargo/appcontext"
)

func init() {
  log.InitLog()
}

type  Indexform struct{
  Name   string
  Password  string
}

type IndexControl struct {
  controller.Controller
}

func (*IndexControl) Before(ctx *appcontext.AppContext,form interface{}) bool {
     log.Debug("Test Before")
     return false
}

func (*IndexControl) Index(ctx *appcontext.AppContext,form interface{}){
     indexform := form.(*Indexform)
     log.Debug(indexform)
}
