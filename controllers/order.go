package controllers

import (
	"encoding/json"
	"fmt"
	"funding/forms"
	"funding/models"
)

// 用户订单相关
type OrderController struct {
	VailUserController
}

// @Title 新增订单
// @Description 新增订单
// @Param	 orderForm		body	forms.NewOrderForm	true	"新增订单信息"
// @Success 200
// @Failure 400
// @router /addOrder [post]
func (c *OrderController) AddOrder() {
	// 获取用户信息
	user := c.User
	//解析 form 表单数据
	var form forms.NewOrderForm
	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	orders, err := models.NewOrderFromForm(user.ID, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(orders)
}

// @Title 根据页面信息来获取订单列表
// @Description	根据页面信息来获取订单列表
// @Param pageForm	body	forms.PageForm	true	"页码信息"
// @Success	200
// @Failure 400
// @router /orderList
func (c *OrderController) OrderList() {
	// 获取用户信息
	user := c.User
	//解析 form 表单数据
	var form forms.NewOrderForm
	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	fmt.Println(user)

	c.ResponseSuccessJson(nil)
}
