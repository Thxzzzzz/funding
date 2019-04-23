package resultModels

// 前端所显示的购物车 item
type CartItem struct {
	ID               uint64 `json:"id"`                 //购物车项ID
	UserId           uint64 `json:"user_id"`            // 用户ID
	ProductPackageId uint64 `json:"product_package_id"` // 套餐ID
	Nums             int    `json:"nums"`               // 购买数量
	Checked          bool   `json:"checked"`            // 是否勾选
	ProductId        uint64 `json:"product_id"`         // 产品ID
	ProductName      string `json:"product_name"`       //产品名称
	Description      string `json:"description"`        //套餐描述
	ImageUrl         string `json:"image_url"`          //套餐图片
}
