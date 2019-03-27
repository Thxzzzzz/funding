package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"testApi/models"
)

type LoginUser struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// 用户相关
type UserControllers struct {
	beego.Controller
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
	c.Data["json"] = result
	c.ServeJSON()
}

// @Title 登录
// @Description 用账号密码登录
// @Param	body	body	controllers.LoginUser	true	""
// @Success 200
// @Failure 400
// @router /login [post]
func (c *UserControllers) Login() {
	//先声明一个 struct 其结构对应请求的 RequestBody 的 Json 结构
	var u LoginUser
	var result models.Result
	//将 RequestBody 的值填充到 struct 之中
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &u)
	//如果解析时出现错误，则说明请求的参数有误
	if err != nil {
		result = models.ErrorResult(models.FALL, err.Error())
	}

	dbResult, err := models.FindUserByUsername(u.Username)

	//数据库查找出错则返回错误
	if err != nil {
		result = models.ErrorResult(models.FALL, err.Error())
	}

	//密码不正确也返回错误
	if dbResult.Password == u.Password {
		result = models.SuccessResult(nil)
		c.SetSession(dbResult.ID, "sdfsdfsdfsd")
		c.Ctx.SetCookie("TestCookie", "Cookie123jldsjflkjsdlkfj")
	} else {
		result = models.ErrorResult(models.FALL, "账号或密码错误")
	}

	c.Data["json"] = result
	c.ServeJSON()
}
