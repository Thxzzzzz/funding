package forms

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
