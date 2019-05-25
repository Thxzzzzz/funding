package managerControllers

import (
	"encoding/json"
	"funding/models"
	"funding/objects"
)

type ManagerOrderController struct {
	VailManagerController
}

// @Title 订单增
// @Description 订单增
// @Param	 form		body	models.Order	true	"订单信息"
// @Success	200
// @Failure	400
// @router /add [post]
func (c *ManagerOrderController) OrderAdd() {
	// 检查是不是超级管理员
	err := c.VerifySuperAdmin()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	var form models.Order
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}

	err = models.InsertOrder(&form)
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}
	c.ResponseSuccessJson(form)
}

// @Title 订单删
// @Description 订单删
// @Param	 form		body	models.Order	true	"订单信息"
// @Success	200
// @Failure	400
// @router /delete [post]
func (c *ManagerOrderController) OrderDelete() {
	// 检查是不是超级管理员
	err := c.VerifySuperAdmin()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	var form models.Order
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}

	err = models.DeleteOrderById(form.ID)
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}
	c.ResponseSuccessJson(form)
}

// @Title 订单查（包括已软删除）
// @Description 订单查（包括已软删除）
// @Param	 form		body	models.Order	true	"订单信息"
// @Success	200
// @Failure	400
// @router /all [get]
func (c *ManagerOrderController) OrderAll() {
	// 检查是不是超级管理员
	err := c.VerifySuperAdmin()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	orders, err := models.FindAllOrderIncludeDeleted()
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}
	c.ResponseSuccessJson(orders)
}

// @Title 订单改（包括已软删除）
// @Description 订单改（包括已软删除）
// @Param	 form		body	models.Order	true	"订单信息"
// @Success	200
// @Failure	400
// @router /update [post]
func (c *ManagerOrderController) OrderUpdate() {
	// 检查是不是超级管理员
	err := c.VerifySuperAdmin()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	var form models.Order
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}
	err = models.UpdateOrderIncludeDeleted(&form)
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}
	c.ResponseSuccessJson(form)
}

// @Title 恢复已删除的订单
// @Description 恢复已删除的订单
// @Param	 form		body	models.Order	true	"订单信息"
// @Success	200
// @Failure	400
// @router /recover [post]
func (c *ManagerOrderController) OrderRecover() {
	// 检查是不是超级管理员
	err := c.VerifySuperAdmin()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	var form models.Order
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}
	err = models.RecoverDeletedOrder(form.ID)
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}
	c.ResponseSuccessJson(form)
}

// @Title 恢复已删除的订单
// @Description 恢复已删除的订单
// @Success	200
// @Failure	400
// @router /complaintOrders [get]
func (c *ManagerOrderController) ComplaintOrders() {
	// 检查是不是超级管理员
	err := c.VerifySuperAdmin()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	result, err := models.GetComplaintOrders()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(result)
}
