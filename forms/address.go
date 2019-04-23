package forms

//地址表单 TODO 表单验证 默认地址
type Address struct {
	ID      uint64 `form:"id"`
	Name    string `form:"name"`    // 收件人姓名
	Address string `form:"address"` // 地址
	Phone   string `form:"phone"`   // 手机号
	Default bool   `form:"default"` // 是否设为默认
}
