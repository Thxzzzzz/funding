package forms

//登录表单
type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// 注册表单 //TODO valid 数据验证
type RegisterForm struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Nickname string `json:"nickname" form:"nickname"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
}

// 用户表
type UserUpdateForm struct {
	ID       uint64 `json:"id"`       // ID
	Username string `json:"username"` // 账号
	Password string `json:"password"` // 密码
	Nickname string `json:"nickname"` // 昵称
	Email    string `json:"email"`    // 邮箱
	Phone    string `json:"phone"`    // 手机号
	IconUrl  string `json:"icon_url"` // 头像
}

// 修改密码表单
type PswForm struct {
	OldPsw string `json:"old_psw"` // 旧密码
	NewPsw string `json:"new_psw"` // 新密码
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
