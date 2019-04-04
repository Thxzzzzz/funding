package forms

//登录表单
type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

//注册表单
type RegisterForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Nickname string `form:"nickname"`
	Email    string `form:"email"`
	Phone    string `form:"phone"`
	IconUrl  string `form:"icon_url"`
}
