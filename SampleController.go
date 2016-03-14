package main

import (
  "github.com/wpxiong/beargo/log"
  "github.com/wpxiong/beargo/controller"
  "github.com/wpxiong/beargo/appcontext"
  "github.com/wpxiong/beargo/render"
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
  log.Debug(sampleForm)
}

func (*SampleControl) Create(ctx *appcontext.AppContext,form interface{}){
  sampleForm := form.(*Sampleform)
  log.Debug(sampleForm)
  log.Debug("Redirect To sample")
  render.RedirectTo(ctx,"/sample/1981?xiong=api&xiong=234")
}
