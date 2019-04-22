package forms

// 购物差表单
type CartForm struct {
	UserId           uint64 `form:"userId"`           // 用户 ID
	ProductPackageId uint64 `form:"productPackageId"` // 套餐 ID
	Nums             int    `form:"nums"`             //数量
	Checked          bool   `form:"checked"`          //是否勾选
}
