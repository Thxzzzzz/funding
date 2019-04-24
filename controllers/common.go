package controllers

import (
	"errors"
	"fmt"
	"funding/models"
	"funding/resultModels"
	"github.com/astaxie/beego"
)

// Session 中保存登录信息的 Key 值
const SESSION_USER_KEY = "userId"

// 定义 Controller 基类
type BaseController struct {
	beego.Controller
}

// 所有请求都需要身份验证的 Controller
type VailUserController struct {
	BaseController
	User *models.User
}

// 实现 Prepare 验证身份
func (c *VailUserController) Prepare() {
	userId := c.GetSession(SESSION_USER_KEY)
	var result *models.User
	if userId == nil {
		c.ResponseErrJson(errors.New("没有登录"))
		return
	}
	id, _ := userId.(uint64)
	// 获取当前 Session 中的 userId 字段对应的值
	result, err := models.FindUserById(id)
	if err != nil {
		c.ResponseErrJson(errors.New("没有该用户"))
		return
	}
	fmt.Println(result)
	c.User = result
}

// 用于返回 err 减少代码量
func (c *BaseController) ResponseErrJson(err error) {
	result := resultModels.ErrorResult(resultModels.FALL, err.Error())
	c.ResponseJson(result)
}

// 返回成功的 Json 数据
func (c *BaseController) ResponseSuccessJson(data interface{}) {
	result := resultModels.SuccessResult(data)
	c.ResponseJson(result)
}

// 用于返回 Json 格式的数据
// 这个项目里会用前后端分离的模式开发，返回的都是 Json 数据
func (c *BaseController) ResponseJson(result resultModels.Result) {
	c.Data["json"] = result
	c.ServeJSON()
}
