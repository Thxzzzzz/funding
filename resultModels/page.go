package resultModels

// 分页信息
type PageInfo struct {
	Page  int `json:"page"`  // 页码
	Total int `json:"total"` // 总数量
}
