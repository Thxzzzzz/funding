package resultModels

import (
	"funding/enums"
	"time"
)

// 订单列表 Item
type OrderListItem struct {
	ID               uint64            `json:"order_id"`           // 订单 Id
	UserId           uint64            `json:"user_id"`            // 买家 Id
	SellerId         uint64            `json:"seller_id"`          // 卖家 Id
	SellerNickname   string            `json:"seller_nickname"`    // 卖家昵称
	ProductId        uint64            `json:"product_id"`         // 产品 Id
	ProductName      string            `json:"product_name"`       // 产品名称
	EndTime          time.Time         `json:"end_time"`           // 众筹截止时间
	ProductPackageId uint64            `json:"product_package_id"` // 套餐 Id
	Stock            int               `json:"stock"`              // 库存
	Description      string            `json:"description"`        // 套餐描述
	ImageUrl         string            `json:"image_url"`          // 图片链接
	Nums             int               `json:"nums"`               // 购买数量
	UnitPrice        float64           `json:"unit_price"`         // 单价
	Freight          float64           `json:"freight"`            // 运费
	TotalPrice       float64           `json:"total_price"`        // 总价
	OrderStatus      enums.OrderStatus `json:"order_status"`       // 订单状态
	CreatedAt        time.Time         `json:"created_at"`         // 创建日期
	Name             string            `json:"name"`               // 收件人姓名
	Phone            string            `json:"phone"`              // 收件人电话
	Address          string            `json:"address"`            // 收件人地址
	PaidAt           *time.Time        `json:"paid_at"`            // 支付时间
	CloseAt          *time.Time        `json:"close_at"`           // 关闭时间
	FinishedAt       *time.Time        `json:"finished_at"`        // 交易成功时间
	//CheckingNumber   string             `json:"checking_number"`    // 物流单号
}

// 订单列表
type OrderList struct {
	PageInfo                   //分页信息
	OrderList []*OrderListItem `json:"order_list"` // 订单列表
}
