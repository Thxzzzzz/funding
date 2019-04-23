package controllers

// 用户订单相关
type OrderController struct {
	VailUserController
}

// @Title 新订单
// @Description 新订单
// @Success 200
// @Failure 400
// @router /newOrder [post]
func (c *OrderController) NewOrder() {
	// TODO 新订单
	c.ResponseSuccessJson(c.User)
}
