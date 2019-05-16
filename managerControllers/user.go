package managerControllers

import (
	"encoding/json"
	"funding/controllers"
	"funding/enums"
	"funding/forms"
	"funding/models"
	"funding/objects"
	"funding/resultModels"
	"github.com/jinzhu/gorm"
)

// 管理端的 UserController
// 一些通用的 API 还是会用 /controllers 里面的
type ManagerUserController struct {
	controllers.BaseController
}

// @Title 注册
// @Description	注册
// @Param user	body	models.User	true	"注册信息"
// @Description 注册
// @Success 200
// @Failure 400
// @router /register [post]
func (c *ManagerUserController) Register() {
	currentUser, err := c.CheckAndGetUser()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 超级管理员才能创建账号
	if currentUser.RoleId != enums.Role_SuperAdmin {
		c.ResponseErrJson(&resultError.UserRoleVerifyError)
		return
	}

	// 首先获取请求中的数据
	//先声明一个 struct 其结构对应请求的 form 表单数据
	form := models.User{}

	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil || form == (models.User{}) {
		c.ResponseErrJson(err)
		return
	}
	//查询是否已存在用户名
	dbExisted, err := models.FindUserByUsername(form.Username)
	//查询出错
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		c.ResponseErrJson(err)
		return
	}
	//已存在
	if dbExisted != nil && dbExisted.Username == form.Username {
		c.ResponseErrJson(resultError.NewFallFundingErr("用户名已存在"))
		return
	}

	//向数据库中插入数据
	err = models.InsertUser(&form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(nil)
}

// @Title 登录
// @Description 用账号密码登录
// @Param	userForm	body	forms.LoginForm		true	"登录账号密码"
// @Success 200 {object} models.User
// @Failure 400
// @router /login [post]
func (c *ManagerUserController) Login() {
	// 1. 首先获取请求中的数据
	//先声明一个 struct 其结构对应请求的 form 表单数据
	loginForm := forms.LoginForm{}
	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &loginForm)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// 2. 获取数据库中的数据并与请求数据进行比较
	dbResult, err := models.FindUserByUsername(loginForm.Username)

	//数据库查找出错则返回错误
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// 这是管理端的登录接口，所以只能是审核员或管理员
	if dbResult.RoleId != enums.Role_Auditor && dbResult.RoleId != enums.Role_SuperAdmin {
		c.ResponseErrJson(&resultError.UserPasswordError)
		return
	}

	// 3. 比较得出结果后，如果正确登录则将信息加入到 Session 中
	if dbResult.Password != loginForm.Password {
		// 密码不正确也返回错误
		c.ResponseErrJson(&resultError.UserPasswordError)
		return
	}

	result := resultModels.SuccessResult(dbResult)
	//向当前 Session 写入 userId
	c.SetSession(controllers.SESSION_USER_KEY, dbResult.ID)

	//  4.. 返回 Json 信息
	c.ResponseJson(result)
}
