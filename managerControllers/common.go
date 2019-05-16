package managerControllers

import (
	"funding/controllers"
	"funding/enums"
	"funding/models"
	"funding/objects"
)

const SESSION_MANAGER_KEY = "managerId"

type ManagerBaseController struct {
	controllers.BaseController
}

//所有请求都需要身份验证的管理者 Controller
type VailManagerController struct {
	ManagerBaseController
	User *models.User
}

func (c *ManagerBaseController) CheckAndGetUser() (*models.User, error) {
	userId := c.GetSession(SESSION_MANAGER_KEY)
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

// 实现 Prepare 验证身份
func (c *VailManagerController) Prepare() {
	user, err := c.CheckAndGetUser()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.User = user
}

// 校验是否为对应角色
func (c *VailManagerController) verifyRole(roleId int) error {
	if c.User.RoleId != roleId {
		return &resultError.UserRoleVerifyError
	}
	return nil
}

// 校验是否是买家
func (c *VailManagerController) VerifyBuyer() error {
	return c.verifyRole(enums.Role_Buyer)
}

// 校验是否是卖家
func (c *VailManagerController) VerifySeller() error {
	return c.verifyRole(enums.Role_Seller)
}

// 校验是否是审核员
func (c *VailManagerController) VerifyAuditor() error {
	return c.verifyRole(enums.Role_Auditor)
}

// 校验是否是审核员
func (c *VailManagerController) VerifySuperAdmin() error {
	return c.verifyRole(enums.Role_SuperAdmin)
}
