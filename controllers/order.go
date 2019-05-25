package controllers

import (
	"encoding/json"
	"funding/enums"
	"funding/forms"
	"funding/models"
	"funding/objects"
	"github.com/astaxie/beego"
	"time"
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
	var orderIds []uint64
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

// @Title 根据订单 ID 确认收货
// @Description  根据订单 ID 确认收货
// @Param	id	body	forms.IdForm	true	"订单ID"
// @Success 200
// @Failure 400
// @router /receivedOrder [Post]
func (c *OrderController) ReceivedOrdergen() {
	user := c.User
	if user.RoleId != enums.Role_Buyer {
		c.ResponseErrJson(resultError.NewFallFundingErr("买家才能确认收货"))
		return
	}
	form := forms.IdForm{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 先根据订单 ID 查询对应订单
	order, err := models.FindOrderById(form.Id)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 订单对应的用户是否是当前请求的用户
	if order.BuyerId != user.ID {
		c.ResponseErrJson(resultError.NewFallFundingErr("这不是你的订单"))
		return
	}
	// 如果不是已发货，则不能确认收货
	if order.Status != enums.OrderStatus_Deliver {
		c.ResponseErrJson(resultError.NewFallFundingErr("订单状态异常"))
		return
	}
	// 已发货则将订单状态改为结束(成功)~
	order.Status = enums.OrderStatus_Finished
	// 并更新订单结束时间
	nowTime := time.Now()
	order.FinishedAt = &nowTime

	// 更新订单信息
	err = models.UpdateOrder(order)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(nil)
}

// @Title 买家根据订单 ID 申请退款
// @Description  买家根据订单 ID 申请退款
// @Param	form	body	forms.RefundForm	true	"订单ID"
// @Success 200
// @Failure 400
// @router /refund [Post]
func (c *OrderController) Refund() {
	user := c.User
	if user.RoleId != enums.Role_Buyer {
		c.ResponseErrJson(resultError.NewFallFundingErr("买家才能申请退款"))
		return
	}
	form := forms.RefundForm{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 先根据订单 ID 查询对应订单
	order, err := models.FindOrderById(form.ID)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 只能给自己的订单申请退款
	if user.RoleId == enums.Role_Buyer && user.ID != order.BuyerId {
		c.ResponseErrJson(resultError.NewFallFundingErr("这不是你的订单"))
		return
	}

	// 如果是已付款且未结束的才能申请退款
	if !(order.Status >= enums.OrderStatus_Paid && order.Status < enums.OrderStatus_Finished) {
		c.ResponseErrJson(resultError.NewFallFundingErr("订单状态异常"))
		return
	}
	// 记录申请退款前的订单状态 (拒绝退款后要恢复到这个状态)
	order.LastStatus = order.Status
	// 修改订单状态为申请退款状态
	order.Status = enums.OrderStatus_Refund
	// 添加退款原因
	order.RefundReason = form.RefundReason
	// 更新订单状态
	err = models.UpdateOrder(order)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(nil)
}

// @Title 买家根据订单 ID 发起投诉
// @Description  买家根据订单 ID 发起投诉
// @Param	form	body	forms.ComplaintForm	true	"订单ID"
// @Success 200
// @Failure 400
// @router /complaint [Post]
func (c *OrderController) Complaint() {
	user := c.User
	if user.RoleId != enums.Role_Buyer {
		c.ResponseErrJson(resultError.NewFallFundingErr("买家才能发起投诉"))
		return
	}
	form := forms.ComplaintForm{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 先根据订单 ID 查询对应订单
	order, err := models.FindOrderById(form.ID)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 只能给自己的订单发起投诉
	if user.RoleId == enums.Role_Buyer && user.ID != order.BuyerId {
		c.ResponseErrJson(resultError.NewFallFundingErr("这不是你的订单"))
		return
	}
	// 添加退款原因
	order.ComplaintReason = form.ComplaintReason
	// 更新订单状态
	err = models.UpdateOrder(order)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(nil)
}

// @Title 根据订单 ID 取消订单
// @Description  根据订单 ID 取消订单
// @Param	form	body	forms.RefundForm	true	"订单ID"
// @Success 200
// @Failure 400
// @router /cancel [Post]
func (c *OrderController) Cancel() {
	user := c.User
	form := forms.RefundForm{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// 先根据订单 ID 查询对应订单
	order, err := models.FindOrderById(form.ID)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// 买家只能取消自己的订单
	if user.RoleId == enums.Role_Buyer && order.BuyerId != user.ID {
		c.ResponseErrJson(resultError.NewFallFundingErr("这不是你的订单"))
		return
	}

	//  买家只能取消未开始备货/发货的，已开始的要先申请退款，卖家同意后取消
	if user.RoleId == enums.Role_Buyer && order.Status >= enums.OrderStatus_Prepare {
		c.ResponseErrJson(resultError.NewFallFundingErr("订单状态异常"))
		return
	}

	// 卖家只能取消卖家为自己的订单
	if user.RoleId == enums.Role_Seller && order.SellerId != user.ID {
		c.ResponseErrJson(resultError.NewFallFundingErr("这不是你的订单"))
		return
	}

	// 取消相关订单
	orderIds := []uint64{form.ID}
	err = models.CancelOrderByOrderIds(orderIds)

	c.ResponseSuccessJson(nil)
}

// @Title 根据订单 ID 删除订单
// @Description  根据订单 ID 取消订单
// @Param	form	body	forms.IdForm	true	"订单ID"
// @Success 200
// @Failure 400
// @router /delOrder [Post]
func (c *OrderController) DelOrder() {
	user := c.User
	form := forms.IdForm{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// 先根据订单 ID 查询对应订单
	order, err := models.FindOrderById(form.Id)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// 买家只能删除自己的订单
	if user.RoleId == enums.Role_Buyer && order.BuyerId != user.ID {
		c.ResponseErrJson(resultError.NewFallFundingErr("这不是你的订单"))
		return
	}

	//  买家只能删除已完成或取消的订单
	if order.Status != enums.OrderStatus_Finished && order.Status != enums.OrderStatus_Canceled {
		c.ResponseErrJson(resultError.NewFallFundingErr("只能删除已完成或取消的订单"))
		return
	}

	// 删除相关订单
	err = models.DeleteOrderById(form.Id)

	c.ResponseSuccessJson(nil)
}

/////////////////// 			商家相关的订单接口					/////////////////

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

// @Title 卖家根据订单 ID 拒绝退款
// @Description   卖家根据订单 ID 拒绝退款
// @Param	form	body	forms.RefundForm	true	"订单ID 和原因"
// @Success 200
// @Failure 400
// @router /cantRefund [Post]
func (c *OrderController) CantRefund() {
	user := c.User
	if user.RoleId != enums.Role_Seller {
		c.ResponseErrJson(resultError.NewFallFundingErr("卖家才能拒绝申请退款"))
		return
	}
	form := forms.RefundForm{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 先根据订单 ID 查询对应订单
	order, err := models.FindOrderById(form.ID)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// 必须是正在申请退款的才能拒绝退款
	if order.Status != enums.OrderStatus_Refund {
		c.ResponseErrJson(resultError.NewFallFundingErr("订单状态异常"))
		return
	}
	// 修改订单状态为申请退款前的状态
	order.Status = order.LastStatus
	// 添加拒绝退款原因
	order.RefundReason = form.RefundReason
	// 更新订单状态
	err = models.UpdateOrder(order)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(nil)
}
