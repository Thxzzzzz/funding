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
	OrderStatus_Refund                                   // 正在退款
	OrderStatus_Canceled                                 //	交易取消
)

// 角色
const (
	Role_Buyer      int = 0   // 买家
	Role_Auditor    int = 1   // 审核员
	Role_Seller     int = 2   // 卖家
	Role_SuperAdmin int = 999 //超级管理员
)

type FundingStatus int

// 众筹状态
const (
	FundingStatus_Success FundingStatus = 1 // 众筹成功
	FundingStatus_Fail    FundingStatus = 2 // 众筹失败
	FundingStatus_Ing     FundingStatus = 3 // 正在众筹
)
