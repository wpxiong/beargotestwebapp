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
<<<<<<< HEAD
  Name   string
  Password  string
=======
  Pam int
>>>>>>> 41fef12e8c9753d2226b34db0458aecf0a0ebe60
}

type IndexControl struct {
  controller.Controller
}

func (*IndexControl) Index(ctx *appcontext.AppContext,form interface{}){
<<<<<<< HEAD
     indexform := form.(*Indexform)
     log.Debug(indexform)
=======
>>>>>>> 41fef12e8c9753d2226b34db0458aecf0a0ebe60
}
