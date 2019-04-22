package resultModels

import "funding/models"

// 前端所显示的购物车 item
type CartItem struct {
	models.Cart
	ProductId   uint64 `json:"product_id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}
