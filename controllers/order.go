package controllers

import (
	"encoding/json"
	"funding/enums"
	"funding/forms"
	"funding/models"
	"funding/objects"
	"github.com/astaxie/beego"
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
	// 不是买家就不能新增订单了吧
	if user.RoleId != enums.Role_Buyer {
		c.ResponseErrJson(resultError.NewFallFundingErr("这不是买家"))
		return
	}
	//解析 form 表单数据
	var form forms.NewOrderForm
	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 返回订单 ID 列表
	orderIdList, err := models.NewOrderFromForm(user.ID, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(orderIdList)
}

// @Title 根据页面信息来获取订单列表
// @Description	根据页面信息来获取订单列表
// @Param pageForm	body	forms.PageForm	true	"页码信息"
// @Success	200
// @Failure 400
// @router /orderList [get]
func (c *OrderController) OrderList() {
	// 获取用户信息
	user := c.User
	//解析 form 表单数据
	var form forms.SellerGetOrderListForm
	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	// 获取所有 query 数据组成的 map
	values := c.Ctx.Request.URL.Query()
	// 解析到 Struct 中
	err := beego.ParseForm(values, &form)
	if err != nil {
		c.ResponseErrJson(err)
	}

	result, err := models.GetOrderListByUserId(&form, user.ID, user.RoleId)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(result)
}

// @Title 根据订单 Id 列表来获取订单列表
// @Description 根据订单 Id 列表来获取订单列表
// @Param	orderId	query	string	true	"订单列表的Json字符串"
// @Success 200
// @Failure 400
// @router /orderInIds [get]
func (c *OrderController) OrderInIds() {
	user := c.User
	// 获取 Json 字符串
	ids := c.GetString("orderId")
	var orderIds []uint64
	err := json.Unmarshal([]byte(ids), &orderIds)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	result, err := models.GetOrderListByOrderIds(orderIds, user.ID, user.RoleId)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(result)
}

// @Title 根据订单 ID 列表来进行支付
// @Description  根据订单 ID 列表来进行支付
// @Param	orderIds	body	[]uint64	true	"订单ID列表"
// @Success 200
// @Failure 400
// @router /orderPay [Post]
func (c *OrderController) OrderPay() {
	orderIds := []uint64{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &orderIds)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	err = models.PayOrderByOrderIdList(orderIds)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(nil)
}

/////////////////// 			商家相关的订单接口					/////////////////

// 这个的 swagger 应该是用不了的，因为懒得把每个字段写进去。。
// @Title 商家获取订单列表
// @Description  商家获取订单列表
// @Param	form	body	forms.SellerGetOrderListForm	true	"订单ID列表"
// @Success 200
// @Failure 400
// @router /orderListToSeller [get]
func (c *OrderController) GetOrderListToSeller() {
	// 校验卖家身份
	err := c.VerifySeller()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 获取用户信息
	user := c.User
	//解析 form 表单数据
	var form forms.SellerGetOrderListForm
	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	// 获取所有 query 数据组成的 map
	values := c.Ctx.Request.URL.Query()
	// 解析到 Struct 中
	err = beego.ParseForm(values, &form)
	if err != nil {
		c.ResponseErrJson(err)
	}
	result, err := models.GetOrderListByUserId(&form, user.ID, user.RoleId)
	if err != nil {
		c.ResponseErrJson(err)
	}
	c.ResponseSuccessJson(result)
}

// @Title 发货
// @Description 发货
// @Param	form	body	forms.OrderSendOutForm	true	"发货参数"
// @Success 200
// @Failure 400
// @router /sendOutOrder [post]
func (c *OrderController) SendOutOrder() {
	// 校验卖家身份
	err := c.VerifySeller()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	form := forms.OrderSendOutForm{}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	err = models.SendOutOrderById(&form, c.User.ID)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(nil)
}
