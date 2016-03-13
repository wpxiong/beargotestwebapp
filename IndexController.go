package main

import (
  "github.com/wpxiong/beargo/log"
  "github.com/wpxiong/beargo/controller"
  "github.com/wpxiong/beargo/appcontext"
  "github.com/wpxiong/beargo/session"
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

func (this *IndexControl) Before(ctx *appcontext.AppContext,form interface{}) bool {
     log.Debug("Test Before")
     return true
}

func (this *IndexControl) Index(ctx *appcontext.AppContext,form interface{}){
     indexform := form.(*Indexform)
     request := ctx.Request.HttpRequest
     response := ctx.Writer.HttpResponseWriter
     var sess session.Session = session.NewSession(request , *response)
     log.Debug("Index Method")
     log.Debug(sess)
     if nil == sess.GetSessionValue("text") {
         sess.SaveSessionValue("test","xiongwenping")
     }
     log.Debug(indexform)
}
