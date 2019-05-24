package forms

import (
	"funding/enums"
)

// 单个订单
type OrderPkgItem struct {
	BuyerId          uint64  `json:"buyer_id"`           // 购买者 ID
	SellerID         uint64  `json:"seller_id"`          // 卖家 ID
	ProductID        uint64  `json:"product_id"`         // 产品 ID
	ProductPackageID uint64  `json:"product_package_id"` // 套餐 ID
	Price            float64 `json:"price"`              // 单价
	Nums             int     `json:"nums"`               // 数量
}

// 新订单表单
type NewOrderForm struct {
	Name         string         `form:"name"`    // 收件人姓名
	Address      string         `form:"address"` // 地址
	Phone        string         `form:"phone"`   // 手机号
	OrderPkgList []OrderPkgItem `json:"order_pkg_list"`
	OrderTotal   float64        `json:"order_total"`
}

// 退款申请表单（带退款原因）
type RefundForm struct {
	ID           uint64 `json:"id"`            // 订单ID
	RefundReason string `json:"refund_reason"` // 退款原因
}

////////////// 			卖家相关								/////////////

// 卖家获取订单的 Form
type SellerGetOrderListForm struct {
	PageForm                          // 页码信息
	OrderStatus   enums.OrderStatus   `json:"order_status" form:"order_status"`     // 订单状态
	FundingStatus enums.FundingStatus `json:"funding_status" form:"funding_status"` // 众筹状态 （0 :全部 ,1:众筹成功,2:失败,3:正在进行）
	ProductId     uint64              `json:"product_id"`                           // 商品ID
}

// 发货请求的表单
type OrderSendOutForm struct {
	OrderId        uint64 `json:"order_id"`        // 订单ID
	CheckingNumber string `json:"checking_number"` // 物流单号
}
