package main

import (
	"funding/controllers"
	_ "funding/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/gomodule/redigo/redis"
)

func main() {
	// 异常处理
	beego.ErrorController(&controllers.ErrorController{})
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
