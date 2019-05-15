package controllers

import (
	"encoding/json"
	"fmt"
	"funding/enums"
	"funding/forms"
	"funding/models"
	"funding/objects"
	"funding/resultModels"
	"funding/utils"
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/gorm"
	"path"
	"time"
)

//type UserValid interface {
//	CheckAndGetUser() *models.User
//}

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
		return nil, &resultError.NotLoginError
	}
	id, ok := userId.(uint64)
	if !ok {
		return nil, &resultError.NotLoginError
	}
	// 获取当前 Session 中的 userId 字段对应的值
	result, err := models.FindUserById(id)
	if err != nil {
		return nil, &resultError.NotLoginError
	}
	return result, nil
}

////////////////			Users	用户信息相关									///////////////

// @Title 根据 id 获取 User
// @Description 根据 Id（数据库表 Id ，不是用户名）来获取对应用户信息
// @Param	id	query	int	true	"数据库 User 表ID"
// @Success 200	{object} models.User
// @Failure 400
// @router /id [get]
func (c *UserControllers) GetUserById() {
	var result resultModels.Result

	//TODO 表单校验，有优化空间啊
	valid := validation.Validation{}
	valid.Required(c.GetString("id"), "id").Message("id 不能为空")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ResponseErrJson(err)
			return
		}
	}

	id, err := c.GetUint64("id")
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	dbResult, err := models.FindUserById(id)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	result = resultModels.SuccessResult(dbResult)
	c.ResponseJson(result)
}

/*
// @Param 	username formData	string 	true	"用户名"
// @Param 	password formData	string 	true	"密码"
// @Param 	nickname formData	string 	true	"昵称"
// @Param 	email    formData	string 	true	"邮箱"
// @Param 	phone    formData	string 	true	"手机号"
*/

// @Title 注册
// @Description	注册
// @Param registerForm	body	forms.RegisterForm	true	"注册信息"
// @Description 注册
// @Success 200
// @Failure 400
// @router /register [post]
func (c *UserControllers) Register() {
	// 首先获取请求中的数据
	//先声明一个 struct 其结构对应请求的 form 表单数据
	form := forms.RegisterForm{}
	var result resultModels.Result

	////将 RequestBody 的值填充到 struct 之中
	//err := c.ParseForm(&form)
	////如果解析时出现错误，则说明请求的参数有误
	//if err != nil {
	//	c.ResponseErrJson(err)
	//	return
	//}

	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil || form == (forms.RegisterForm{}) {
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
		result = resultModels.ErrorResult(resultModels.FALL, "用户名已存在")
		c.ResponseJson(result)
		return
	}

	user := models.User{}
	err = utils.CopyStruct(form, &user)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	//向数据库中插入数据
	err = models.InsertUser(&user)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	//注册成功后将 id 加入到 session 中,即可记录登录状态
	c.SetSession(SESSION_USER_KEY, user.ID)
	result = resultModels.SuccessResult(nil)
	c.ResponseJson(result)
}

// @Title 登录
// @Description 用账号密码登录
// @Param	userForm	body	forms.LoginForm		true	"登录账号密码"
// @Success 200 {object} models.User
// @Failure 400
// @router /login [post]
func (c *UserControllers) Login() {
	// 1. 首先获取请求中的数据
	//先声明一个 struct 其结构对应请求的 form 表单数据
	loginForm := forms.LoginForm{}
	//var result resultModels.Result

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

	// 这是商城端的登录接口，所以只能是商家或者买家登录
	if dbResult.RoleId != enums.Role_Buyer && dbResult.RoleId != enums.Role_Seller {
		c.ResponseErrJson(&resultError.UserNotExitError)
		return
	}

	// 3. 比较得出结果后，如果正确登录则将信息加入到 Session 中
	if dbResult.Password != loginForm.Password {
		// 密码不正确也返回错误
		c.ResponseErrJson(&resultError.UserNotExitError)
		return
	}

	result := resultModels.SuccessResult(dbResult)
	//向当前 Session 写入 userId
	c.SetSession(SESSION_USER_KEY, dbResult.ID)
	//TODO 单点登录

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
// @Success 200	{object} models.User
// @Failure 400
// @router /info [get]
func (c *UserControllers) Info() {
	var result resultModels.Result
	user, err := c.CheckAndGetUser()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	result = resultModels.SuccessResult(user)
	c.ResponseJson(result)
}

const imagesPath = "uploadfile/images"
const baseUrl = "http://127.0.0.1:8080/"

// @Title 上传图片
// @Description 上传图片
// @Param	file	formData	multipart.File	true	"图片文件"
// @Accept form
// @Success 200
// @Failure 400
// @router /uploadImage [post]
func (c *UserControllers) UploadImage() {
	_, err := c.CheckAndGetUser()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	file, information, err := c.GetFile("file") //返回文件，文件信息头，错误信息
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	defer file.Close()                                                      //关闭上传的文件，否则出现临时文件不清除的情况
	filename := fmt.Sprintf("%d", time.Now().Unix()) + information.Filename //将文件信息头的信息赋值给filename变量
	imgPath := path.Join(imagesPath, filename)                              //图片保存地址
	err = c.SaveToFile("file", imgPath)                                     //保存文件的路径。保存在static/upload中   （文件名）
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(baseUrl + imgPath)
}

// 仅仅是为了配合 element-ui 的上传空间写的 options 请求 api,返回成功即可
// @router /uploadImage [options]
func (c *UserControllers) OptionsUploadImage() {
	c.ResponseSuccessJson(nil)
}
