package forms

// 请求评论信息表单
type CommentListByProductForm struct {
	ProductId uint64 `json:"product_id" form:"product_id"`
	PageForm
}
