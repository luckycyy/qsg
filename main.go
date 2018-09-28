package main

import (
	_ "qsg/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/qiseguang?charset=utf8&loc=Asia%2FShanghai")
	orm.Debug = true
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}

