package resultModels

import "time"

// 前端所显示的购物车 item
type CartItem struct {
	ID               uint64    `json:"id"`                 // 购物车项ID
	UserId           uint64    `json:"user_id"`            // 用户ID
	SellerId         uint64    `json:"seller_id"`          // 卖家ID
	ProductPackageId uint64    `json:"product_package_id"` // 套餐ID
	Price            float64   `json:"price"`              // 单价
	Stock            int       `json:"stock"`              // 库存
	Nums             int       `json:"nums"`               // 购买数量
	Checked          bool      `json:"checked"`            // 是否勾选
	ProductId        uint64    `json:"product_id"`         // 产品ID
	ProductName      string    `json:"product_name"`       // 产品名称
	Description      string    `json:"description"`        // 套餐描述
	ImageUrl         string    `json:"image_url"`          // 套餐图片
	EndTime          time.Time `json:"end_time"`           // 截止时间
}
