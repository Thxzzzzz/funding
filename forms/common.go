package forms

type IdForm struct {
	Id uint64 `form:"id"`
}

// 页码信息
type PageForm struct {
	Page     int `form:"page"`      // 页码
	PageSize int `form:"page_size"` // 每页数量
}

// pagination: {
//   currentPage: 1,
//   pageSize: 5,
//   total: 0
// },

// 管理系统的页码信息
type PaginationForm struct {
	CurrentPage int `json:"currentPage" form:"currentPage"`
	PageSize    int `json:"pageSize" form:"pageSize"`
	//Total       int `json:"total" form:"total"`
}
