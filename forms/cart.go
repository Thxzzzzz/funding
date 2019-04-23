package forms

// 购物车表单
type CartForm struct {
	UserId           uint64 `form:"user_id"`            // 用户 ID
	ProductPackageId uint64 `form:"product_package_id"` // 套餐 ID
	Nums             int    `form:"nums"`               //数量
	Checked          bool   `form:"checked"`            //是否勾选
}

// 删除购物车 item 表单
type DelCartForm struct {
	ProductPackageId uint64 `form:"product_package_id"` // 套餐 ID
}
