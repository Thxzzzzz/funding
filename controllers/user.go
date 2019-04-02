package controllers

import (
	"testApi/forms"
	"testApi/models"
)

// 用户相关
type UserControllers struct {
	//嵌入 BaseController
	BaseController
}

// @Title 根据 id 获取 User
// @Description 根据 Id（数据库表 Id ，不是用户名）来获取对应用户信息
// @Param	id	path	int	true	"数据库 User 表ID"
// @Success 200
// @Failure 400
// @router /id/:id [get]
func (c *UserControllers) GetUserById() {
	// 这里的 Key 要注意带上冒号，否则获取不到对应的参数
	idd, err := c.GetInt64(":id")
	dbResult, err := models.FindUserById(idd)
	var result models.Result
	if err != nil {
		result = models.ErrorResult(models.FALL, err.Error())
	} else {
		result = models.SuccessResult(dbResult)
	}
	c.ResponseJson(result)
}

// Session 中保存登录信息的 Key 值
const SESSION_USER_KEY = "userId"

// @Title 登录
// @Description 用账号密码登录
// @Param	LoginUserForm	body	forms.LoginForm		true	"登录信息"
// @Success 200
// @Failure 400
// @router /login [post]
func (c *UserControllers) Login() {
	// 1. 首先获取请求中的数据
	//先声明一个 struct 其结构对应请求的 RequestBody 的 Json 结构
	loginForm := forms.LoginForm{}
	var result models.Result
	//将 RequestBody 的值填充到 struct 之中
	err := c.ParseForm(&loginForm)
	//如果解析时出现错误，则说明请求的参数有误
	if err != nil {
		result = models.ErrorResult(models.FALL, err.Error())
	}

	// 2. 获取数据库中的数据并与请求数据进行比较
	dbResult, err := models.FindUserByUsername(loginForm.Username)

	//数据库查找出错则返回错误
	if err != nil {
		result = models.ErrorResult(models.FALL, err.Error())
	}

	// 3. 比较得出结果后，如果正确登录则将信息加入到 Session 中
	if dbResult.Password == loginForm.Password {
		result = models.SuccessResult(nil)
		//向当前 Session 写入 userId
		c.SetSession(SESSION_USER_KEY, dbResult.ID)
		//TODO 单点登录
	} else {
		// 密码不正确也返回错误
		result = models.ErrorResult(models.FALL, "用户名或密码错误")
	}
	//  4.. 返回 Json 信息
	c.ResponseJson(result)
}

//	@Title 登出
//	@router /logout	[post]
func (c *UserControllers) Logout() {
	var result models.Result
	//直接销毁 Session
	c.DestroySession()
	result = models.SuccessResult(nil)
	c.ResponseJson(result)
}

//	@Title 根据当前的 Session 查询对应的用户信息
//	@router /info	[get]
func (c *UserControllers) Info() {
	userId := c.GetSession(SESSION_USER_KEY)
	var result models.Result
	if userId == nil {
		result = models.ErrorResult(models.FALL, "没有登录")
	} else {
		// 获取当前 Session 中的 userId 字段对应的值
		user, err := models.FindUserById(int64(userId.(uint)))
		if err != nil {
			result = models.ErrorResult(models.FALL, "没有该用户")
		} else {
			result = models.SuccessResult(user)
		}
	}
	c.ResponseJson(result)
}
