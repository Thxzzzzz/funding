package managerControllers

import (
	"encoding/json"
	"funding/enums"
	"funding/forms"
	"funding/models"
	"funding/objects"
	"funding/resultModels"
	"funding/utils"
	"github.com/jinzhu/gorm"
)

// 管理端的 UserController
type ManagerUserController struct {
	ManagerBaseController
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
	c.SetSession(SESSION_MANAGER_KEY, dbResult.ID)

	//  4.. 返回 Json 信息
	c.ResponseJson(result)
}

// @Title 注销
// @Description 注销
// @Success 200
// @Failure 400
// @router /logout [post]
func (c *ManagerUserController) Logout() {
	var result resultModels.Result
	//直接销毁 Session
	c.DestroySession()
	result = resultModels.SuccessResult(nil)
	c.ResponseJson(result)
}

// @Title 根据当前的 Session 查询对应的用户信息
// @Description 用户登录之后查询当前登录的用户信息，每次查询会刷新 Session 有效期
// @Success 200	{object} models.User
// @Failure 400
// @router /info [get]
func (c *ManagerUserController) Info() {
	var result resultModels.Result
	user, err := c.CheckAndGetUser()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	result = resultModels.SuccessResult(user)
	c.ResponseJson(result)
}

// @Title 根据角色ID来获取用户列表
// @Description 根据角色ID来获取用户列表
// @Param	role_id	query	int	true	"角色ID"
// @Success 200	{object} models.User
// @Failure 400
// @router /infoByRoleId [get]
func (c *ManagerUserController) GetInfoByRoleId() {
	user, err := c.CheckAndGetUser()
	if err != nil || user.RoleId != enums.Role_SuperAdmin {
		c.ResponseErrJson(&resultError.UserRoleVerifyError)
		return
	}
	roleId, err := c.GetInt("role_id")
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 根据 role_id 来查询角色列表
	result, err := models.GetUsersByRoleId(roleId)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(result)
}

// @Title 更新用户信息
// @Description 更新用户信息
// @Param	form	body	forms.UserFormWithRole	true	"用户信息"
// @Success 200	{object} models.User
// @Failure 400
// @router /updateUser [post]
func (c *ManagerUserController) UpdateUser() {
	user, err := c.CheckAndGetUser()
	if err != nil {
		c.ResponseErrJson(&resultError.UserRoleVerifyError)
		return
	}
	form := forms.UserFormWithRole{}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 如果不是超级管理员，那只能更新自己的信息
	if user.RoleId != enums.Role_SuperAdmin && user.ID != form.ID {
		c.ResponseErrJson(&resultError.UserRoleVerifyError)
		return
	}

	//因为 User 屏蔽了 password 的 json 解析，所以这里要转一下
	newUser := models.User{}
	err = utils.CopyStruct(form, &newUser)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// 更新信息
	err = models.UpdateUser(&newUser)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(newUser)
}

// @Title 更新用户信息
// @Description 更新用户信息
// @Param	form	body	forms.UserFormWithRole	true	"用户信息"
// @Success 200	{object} models.User
// @Failure 400
// @router /newUser [post]
func (c *ManagerUserController) NewUser() {
	user, err := c.CheckAndGetUser()
	// 超级管理员才能创建
	if err != nil || user.RoleId != enums.Role_SuperAdmin {
		c.ResponseErrJson(&resultError.UserRoleVerifyError)
		return
	}
	form := forms.UserFormWithRole{}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
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
		c.ResponseErrJson(resultError.NewFallFundingErr("用户已存在"))
		return
	}

	//因为 User 屏蔽了 password 的 json 解析，所以这里要转一下
	newUser := models.User{}
	err = utils.CopyStruct(form, &newUser)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// 更新信息
	err = models.InsertUser(&newUser)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(newUser)
}
