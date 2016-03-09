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

type Person struct {
    Name string
    Age int
}

type AddressInfo struct {
     StreetAddress string
     City   string
     State string
     PostalCode  string
}

type PhoneNumberInfo struct {
   Type string
   Number string

}


type  Sampleform struct{
  Pam  int
  Pam2  uint8
  Pam3  float32
  Test  string
  Xiong  []int
  PersonInfo Person
  FirstName string
  LastName string 
  Age int
  Address AddressInfo 
  PhoneNumber []PhoneNumberInfo
  
}


type SampleControl struct {
  controller.Controller
}

func (*SampleControl) Index(ctx *appcontext.AppContext,form interface{}){
  sampleForm := form.(*Sampleform)
  fmt.Fprint(*(ctx.Writer.HttpResponseWriter),sampleForm)
}

func (*SampleControl) Create(ctx *appcontext.AppContext,form interface{}){
  sampleForm := form.(*Sampleform) 
  fmt.Fprint(*(ctx.Writer.HttpResponseWriter),sampleForm)
}
