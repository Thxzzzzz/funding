package forms

// 获取产品列表的表单
type ProductListForm struct {
	Page          int    `form:"page"`           // 页码
	PageSize      int    `form:"page_size"`      // 每页数量
	Name          string `form:"name"`           // 产品名称
	Type          int    `form:"type"`           // 产品类型
	Sort          string `form:"sort"`           // 按照 products 中哪个字段排序
	Asc           bool   `form:"asc"`            // 升序
	FundingStatus int    `form:"funding_status"` // 众筹状态
	//PriceGt       float64 `form:"price_gt"`       // 价格大于
	//PriceLte      float64 `form:"price_lte"`      // 价格小于
}
