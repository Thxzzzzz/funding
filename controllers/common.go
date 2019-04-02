package controllers

import (
	"github.com/astaxie/beego"
	"testApi/models"
)

//定义 Controller 基类
type BaseController struct {
	beego.Controller
}

func (c *BaseController) ResponseJson(result models.Result) {
	c.Data["json"] = result
	c.ServeJSON()
}
