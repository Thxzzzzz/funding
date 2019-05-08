package enums

//审核状态枚举
const (
	Verify_Fail    = 0 //未通过
	Verify_Success = 1 //已通过
	Verify_Wait    = 2 //未审核
)

// 订单状态
type OrderStatus int

// 订单状态
const (
	OrderStatus_Ordered  OrderStatus = OrderStatus(iota) //	下单
	OrderStatus_Paid                                     //	已支付
	OrderStatus_Prepare                                  //	正在配货
	OrderStatus_Deliver                                  //	出货 配送
	OrderStatus_Finished                                 //	交易成功
	OrderStatus_Refund                                   // 申请退款
	OrderStatus_Canceled                                 //	交易取消
)
