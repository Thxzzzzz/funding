package controllers

import (
	"errors"
	"fmt"
	"funding/forms"
	"funding/models"
	"funding/resultModels"
	"funding/utils"
)

// 用户相关
type UserControllers struct {
	//嵌入 BaseController
	BaseController
}

// 检查并获取对应的 User 信息
func (c *UserControllers) CheckAndGetUser() (*models.User, error) {
	userId := c.GetSession(SESSION_USER_KEY)
	var result *models.User
	if userId == nil {
		return nil, errors.New("没有登录")
	}
	id, _ := userId.(uint64)
	// 获取当前 Session 中的 userId 字段对应的值
	result, err := models.FindUserById(id)
	if err != nil {
		return nil, errors.New("没有该用户")
	}
	return result, nil
}

// @Title 根据 id 获取 User
// @Description 根据 Id（数据库表 Id ，不是用户名）来获取对应用户信息
// @Param	id	query	int	true	"数据库 User 表ID"
// @Success 200
// @Failure 400
// @router /id [get]
func (c *UserControllers) GetUserById() {
	// 这里的 Key 要注意带上冒号，否则获取不到对应的参数
	idd, err := c.GetUint64("id")
	dbResult, err := models.FindUserById(idd)
	var result resultModels.Result
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
	} else {
		result = resultModels.SuccessResult(dbResult)
	}
	c.ResponseJson(result)
}

// @Title 注册
// @Description 注册
// @Param 	username formData	string 	true	"用户名"
// @Param 	password formData	string 	true	"密码"
// @Param 	nickname formData	string 	true	"昵称"
// @Param 	email    formData	string 	true	"邮箱"
// @Param 	phone    formData	string 	true	"手机号"
// @Success 200
// @Failure 400
// @router /register [post]
func (c *UserControllers) Register() {
	// 首先获取请求中的数据
	//先声明一个 struct 其结构对应请求的 form 表单数据
	form := forms.RegisterForm{}
	var result resultModels.Result
	//将 RequestBody 的值填充到 struct 之中
	err := c.ParseForm(&form)
	//如果解析时出现错误，则说明请求的参数有误
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
		c.ResponseJson(result)
		return
	}

	//查询是否已存在用户名
	dbExisted, err := models.FindUserByUsername(form.Username)
	//查询出错
	if err != nil && err.Error() != "record not found" {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
		c.ResponseJson(result)
		return
	}
	//已存在
	if dbExisted != nil && dbExisted.Username == form.Username {
		result = resultModels.ErrorResult(resultModels.FALL, "用户名已存在")
		c.ResponseJson(result)
		return
	}

	user := models.User{}
	err = utils.CopyStruct(form, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	//向数据库中插入数据
	err = models.InsertUser(&user)
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
		c.ResponseJson(result)
		return
	}
	//注册成功后将 id 加入到 session 中,即可记录登录状态
	c.SetSession(SESSION_USER_KEY, user.ID)
	result = resultModels.SuccessResult(nil)
	c.ResponseJson(result)
}

// @Title 登录
// @Description 用账号密码登录
// @Param	username	formData	string		true	"用户名"
// @Param	password	formData	string		true	"密码"
// @Success 200
// @Failure 400
// @router /login [post]
func (c *UserControllers) Login() {
	// 1. 首先获取请求中的数据
	//先声明一个 struct 其结构对应请求的 form 表单数据
	loginForm := forms.LoginForm{}
	var result resultModels.Result
	//将 RequestBody 的值填充到 struct 之中
	err := c.ParseForm(&loginForm)
	//如果解析时出现错误，则说明请求的参数有误
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
	}

	// 2. 获取数据库中的数据并与请求数据进行比较
	dbResult, err := models.FindUserByUsername(loginForm.Username)

	//数据库查找出错则返回错误
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
	}

	// 3. 比较得出结果后，如果正确登录则将信息加入到 Session 中
	if dbResult.Password == loginForm.Password {
		result = resultModels.SuccessResult(nil)
		//向当前 Session 写入 userId
		c.SetSession(SESSION_USER_KEY, dbResult.ID)
		//TODO 单点登录
	} else {
		// 密码不正确也返回错误
		result = resultModels.ErrorResult(resultModels.FALL, "用户名或密码错误")
	}
	//  4.. 返回 Json 信息
	c.ResponseJson(result)
}

// @Title 注销
// @Description 注销
// @Success 200
// @Failure 400
// @router /logout [post]
func (c *UserControllers) Logout() {
	var result resultModels.Result
	//直接销毁 Session
	c.DestroySession()
	result = resultModels.SuccessResult(nil)
	c.ResponseJson(result)
}

// @Title 根据当前的 Session 查询对应的用户信息
// @Description 用户登录之后查询当前登录的用户信息，每次查询会刷新 Session 有效期
// @Success 200
// @Failure 400
// @router /info [get]
func (c *UserControllers) Info() {
	var result resultModels.Result
	user, err := c.CheckAndGetUser()
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
		c.ResponseJson(result)
		return
	}
	result = resultModels.SuccessResult(user)
	c.ResponseJson(result)
}

// @Title 根据 userId 获取收货地址
// @Description 根据 userId 获取收货地址
// @Success 200	{object} []models.Address
// @Failure 400
// @router /addresses [get]
func (c *UserControllers) GetAddresses() {
	var result resultModels.Result
	user, err := c.CheckAndGetUser()
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
		c.ResponseJson(result)
		return
	}
	addresses, err := models.FindAddressesByUserId(user.ID)
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
		c.ResponseJson(result)
		return
	}
	result = resultModels.SuccessResult(addresses)
	c.ResponseJson(result)
}

// @Title 'TODO:添加购物车'
// @Description 添加购物车
// @router /addCart [post]
func (c *UserControllers) AddCart() {

}
