package main

import (
  "github.com/wpxiong/beargo/log"
  "github.com/wpxiong/beargo/webapp"
  "github.com/wpxiong/beargo/appcontext"
  "github.com/wpxiong/beargo/moudle"
  "time"
  "runtime"
)



type AddressInformation struct {
  Addressid int `id:"true"`
  Userid int
  AddressName string  `notnull:"true"     length:"128" `
}

type ClassInfo struct {
  ClassName string `length:"128"`
  ClassId int `id:"true"   auto_increment:"true"`
  Address  string `length:"128"  notnull:"true"`
  Goup    []GroupInfo `relation_type:"onetomany"  column_name:"classid" referenced_column_name:"classid"`
}


type GroupInfo struct {
  UserId  int  `id:"true"`
  ClassId int  `id:"true"`
  Class ClassInfo  `relation_type:"manytoone" column_name:"classid"  referenced_column_name:"classid"`
  User  UserInfo  `relation_type:"manytoone"   column_name:"userid" referenced_column_name:"id"`
}


type UserInfo struct {
  Id      int `id:"true"`
  UserName string   `id:"true"     notnull:"true"     length:"128"     default_value:"testping"`
  UserAge  uint16  `default_value:"54"`
  UserSex  bool  `default_value:"true"`
  Test     complex64 `default_value:"12.34,11.78"`
  CreateTime  time.Time
  Addressid  int
  Goup    []GroupInfo `relation_type:"onetomany" column_name:"id" referenced_column_name:"userid"`
  Address  AddressInformation `relation_type:"onetoone" column_name:"addressid" referenced_column_name:"addressid"`
}


func main() {
   log.InitLogWithLevel("Debug")
   configMap := InitConfig()
   
   moudleInstance :=  moudle.CreateModuleInstance(moudle.MYSQL,"test","tcp(localhost:3306)","","") 
   
   moudleInstance.AddTable(UserInfo{})
   moudleInstance.AddTable(ClassInfo{})
   moudleInstance.AddTable(GroupInfo{})
   moudleInstance.AddTable(AddressInformation{})

   moudleInstance.InitialDB(true)
   var inserInfo *moudle.DbSqlBuilder = moudleInstance.Insert(AddressInformation{Addressid:1,Userid:12,AddressName:"tokyo shibuya"})
   inserInfo.InsertExecute()
   userinfo := UserInfo{Addressid:1,UserName:"xiong",UserAge:22,UserSex:true,Test:complex(23.4,56.7),CreateTime:time.Now(),Id:1}
   
   moudleInstance.Insert(AddressInformation{Addressid:2,Userid:12,AddressName:"tokyo shibuya"}).InsertExecute()
   moudleInstance.Insert(userinfo).InsertExecute()
   userinfo.UserName = "baili"
   userinfo.Addressid = 2
   userinfo.Id = 2
   moudleInstance.Insert(userinfo).InsertExecute()
   
   moudleInstance.Insert(ClassInfo{ClassName:"class1",Address:"address1"}).InsertExecute()
   moudleInstance.Insert(ClassInfo{ClassName:"class2",Address:"address2"}).InsertExecute()
   
   userinfo.UserAge = 45
   userinfo.UserSex = false
   userinfo.UserName = "test"
   moudleInstance.Update(userinfo)
   userinfo.UserAge = 46
   userinfo.Test = complex(23.4,563.7)
   
   moudleInstance.UpdateWithField(userinfo,[]string{"Test"})
   moudleInstance.UpdateWithWhere(userinfo,[]string{"UserSex"}).WhereOr([]string{"Id","UserName"}).SaveExecute()
   
   moudleInstance.Insert(GroupInfo{UserId:1,ClassId:1}).InsertExecute()
   moudleInstance.Insert(GroupInfo{UserId:1,ClassId:2}).InsertExecute()
   moudleInstance.Insert(GroupInfo{UserId:2,ClassId:1}).InsertExecute()
   moudleInstance.Insert(GroupInfo{UserId:2,ClassId:2}).InsertExecute()
   
   var info *moudle.DbSqlBuilder  = moudleInstance.Query(UserInfo{},moudle.EAGER,[]string{"Goup","Address"})
   //var info moudle.DbSqlBuilder  = moudleInstance.SimpleQuery(GroupInfo{})
   list := info.FetchAll()
   for _,val := range list {
      usr := val.(UserInfo)
      log.Debug(usr)
   }
   
   var ts *moudle.Trans = moudleInstance.Begin()
   
   ts.Delete(GroupInfo{UserId:2,ClassId:2})
   pam := make([]interface{},2)
   pam[0] = 1
   pam[1] = 2
   ts.DeleteWithSql("delete from groupinfo where userid = ? and classid = ?" , pam)
   
   ts.Rollback()
   
   runtime.GOMAXPROCS(runtime.NumCPU())
   config := appcontext.AppConfigContext{Port :9001,ConfigPath : "./setting.conf"}
   
   var appCon appcontext.AppContext = appcontext.AppContext{ ConfigContext :  &config}
   
   app := webapp.New(&appCon,configMap)
   
   indexCtrl := &IndexControl{}
   sampleCtrl := &SampleControl{}
   app.AddAutoRoute("/<name:.*>/<password:[0-9]+>",indexCtrl,Indexform{})
   app.AddAutoRouteWithViewPath("/sample/<pam:int>/",sampleCtrl,Sampleform{},"/sampletest")
   
   app.Start()
}