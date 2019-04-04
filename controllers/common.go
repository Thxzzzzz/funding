package controllers

import (
	"funding/models"
	"github.com/astaxie/beego"
)

// Session 中保存登录信息的 Key 值
const SESSION_USER_KEY = "userId"

//定义 Controller 基类
type BaseController struct {
	beego.Controller
}

// 用于返回 Json 格式的数据
// 这个项目里会用前后端分离的模式开发，返回的都是 Json 数据
func (c *BaseController) ResponseJson(result models.Result) {
	c.Data["json"] = result
	c.ServeJSON()
}
