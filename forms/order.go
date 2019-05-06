package forms

// 单个订单
type OrderPkgItem struct {
	UserID           uint64  `json:"user_id"`            // 购买者 ID
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
