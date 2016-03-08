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
  Pam int
}

type IndexControl struct {
  controller.Controller
}

func (*IndexControl) Index(ctx *appcontext.AppContext,form interface{}){
}
