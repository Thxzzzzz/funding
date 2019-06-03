package main

import (
	"funding/controllers"
	_ "funding/routers"
	"funding/task"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/gomodule/redigo/redis"
)

func main() {
	// 异常处理
	beego.ErrorController(&controllers.ErrorController{})
	// 将 uploadfile 作为静态资源路径
	beego.SetStaticPath("/uploadfile", "uploadfile")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	task.StartTask()
	defer task.StopTask()

	beego.Run()
}
