package forms

type IdForm struct {
	Id uint64 `form:"id"`
}

// 页码信息
type PageForm struct {
	Page     int `form:"page"`      // 页码
	PageSize int `form:"page_size"` // 每页数量
}
