package enums

//审核状态枚举
const (
	VerifyFail    = 0 //未通过
	VerifySuccess = 1 //已通过
	VerifyWait    = 2 //未审核
)

type OrderStatus int

const (
	Ordered  OrderStatus = OrderStatus(iota) //	下单
	Paid                                     //	支付
	Prepare                                  //	配货
	Deliver                                  //	出货 配送
	Finished                                 //	交易成功
	Refund                                   // 退款
	Canceled                                 //	交易失败
)
