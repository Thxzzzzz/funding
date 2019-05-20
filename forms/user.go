package forms

//登录表单
type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

//注册表单 //TODO valid 数据验证
type RegisterForm struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Nickname string `json:"nickname" form:"nickname"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
}

//注册表单 带角色
type UserFormWithRole struct {
	ID       uint64 `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Nickname string `json:"nickname" form:"nickname"`
	RoleId   int    `json:"role_id"  gorm:"default:0"` //角色  ( 0:普通用户(默认) )
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
}
