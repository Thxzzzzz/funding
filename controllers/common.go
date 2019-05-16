package controllers

import (
	"funding/enums"
	"funding/models"
	"funding/objects"
	"funding/resultModels"
	"github.com/astaxie/beego"
)

// Session 中保存登录信息的 Key 值
const SESSION_USER_KEY = "userId"

// 定义 Controller 基类
type BaseController struct {
	beego.Controller
}

// 检查并获取对应的 User 信息
func (c *BaseController) CheckAndGetUser() (*models.User, error) {
	userId := c.GetSession(SESSION_USER_KEY)
	var result *models.User
	if userId == nil {
		return nil, &resultError.NotLoginError
	}
	id, ok := userId.(uint64)
	if !ok {
		return nil, &resultError.NotLoginError
	}
	// 获取当前 Session 中的 userId 字段对应的值
	result, err := models.FindUserById(id)
	if err != nil {
		return nil, &resultError.UserNotExitError
	}
	return result, nil
}

// 所有请求都需要身份验证的 Controller
type VailUserController struct {
	BaseController
	User *models.User
}

// 实现 Prepare 验证身份
func (c *VailUserController) Prepare() {
	user, err := c.CheckAndGetUser()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.User = user
}

// 校验是否为对应角色
func (c *VailUserController) verifyRole(roleId int) error {
	if c.User.RoleId != roleId {
		return &resultError.UserRoleVerifyError
	}
	return nil
}

// 校验是否是买家
func (c *VailUserController) VerifyBuyer() error {
	return c.verifyRole(enums.Role_Buyer)
}

// 校验是否是卖家
func (c *VailUserController) VerifySeller() error {
	return c.verifyRole(enums.Role_Seller)
}

// 校验是否是审核员
func (c *VailUserController) VerifyAuditor() error {
	return c.verifyRole(enums.Role_Auditor)
}

// 校验是否是审核员
func (c *VailUserController) VerifySuperAdmin() error {
	return c.verifyRole(enums.Role_SuperAdmin)
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
