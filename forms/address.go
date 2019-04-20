package forms

//地址表单 TODO 表单验证 默认地址
type Address struct {
	ID      uint64 `form:"id"`
	Name    string `form:"name"`
	Address string `form:"address"`
	Phone   string `form:"phone"`
	Default bool   `form:"default"` //是否设为默认
}
