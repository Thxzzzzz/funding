package controllers

import (
	"encoding/json"
	"errors"
	"funding/forms"
	"funding/models"
	"funding/resultModels"
	"funding/utils"
)

// 用户地址相关
type AddressController struct {
	VailUserController
}

////////// 					 Address 收货地址相关								///////////

// @Title 根据 userId 获取收货地址
// @Description 根据 userId 获取收货地址
// @Success 200	{object} []models.Address
// @Failure 400
// @router /all [get]
func (c *AddressController) GetAddresses() {
	// 获取用户信息
	user := c.User
	//根据 UserId 来查找对应的地址
	addresses, err := models.FindAddressesByUserId(user.ID)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	result := resultModels.SuccessResult(addresses)
	c.ResponseJson(result)
}

// @Title 添加新的地址
// @Description	添加新的地址
// @Param	addressForm	body	forms.Address	true	"地址表单"
// @Success	200
// @Failure 400
// @router /new [post]
func (c *AddressController) NewAddress() {

	// @Param	name	formData	string	true	"收货人姓名"
	// @Param	address	formData	string	true	"收货地址"
	// @Param	phone	formData	string	true	"联系电话"
	// @Param	default	formData	bool	false	"默认地址"

	// 获取用户信息
	user := c.User

	//解析 form 表单数据
	var form forms.Address
	//err = c.ParseForm(&form)
	//if err != nil {
	//	c.ResponseErrJson(err)
	//	return
	//}

	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	//将表单数据复制到 Address 中
	address := models.Address{}
	err = utils.CopyStruct(form, &address)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	//设置 UserId
	address.UserId = user.ID
	err = models.InsertAddress(&address)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	if form.Default {
		user.DefaultAddressId = address.ID
		err = models.UpdateUser(user)
		if err != nil {
			c.ResponseErrJson(errors.New("未能成功设置为默认地址"))
			return
		}
	}
	c.ResponseSuccessJson(nil)

	//result := resultModels.SuccessResult(nil)
	//c.ResponseJson(result)
}

// @Title 根据地址的 id 删除对应的地址
// @Description	根据地址的 id 删除对应的地址
// @Param	id	formData	string	true	"地址的id"
// @Success 200
// @Failure 400
// @router	/delete [post]
func (c *AddressController) DeleteAddress() {

	// 获取用户信息
	user := c.User

	//首先要检查登录状态
	var result resultModels.Result

	var form forms.IdForm
	//TODO 或许应该表单校验
	//aId, err := c.GetUint64("id")

	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	//根据请求的 id 查找对应地址
	address, err := models.FindAddressById(form.Id)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// ID 不匹配 返回错误信息
	if address.UserId != user.ID {
		result = resultModels.ErrorResult(resultModels.FALL, "记录与用户不匹配，这不是你的地址~")
		c.ResponseJson(result)
		return
	}
	// ID 匹配，删除对应的数据
	err = models.DeleteAddressById(form.Id)
	// 删除出错 返回错误信息
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	//删除成功，返回成功提示
	result = resultModels.SuccessResult(nil)
	c.ResponseJson(result)
}

// @Title 更新指定 id 的地址
// @Description	添加新的地址
// @Param	addressForm	body	forms.Address	true	"地址表单"
// @Success	200
// @Failure 400
// @router /update [post]
func (c *AddressController) UpdateAddress() {

	// @Param	id		formData	string	true	"地址ID"
	// @Param	name	formData	string	false	"收货人姓名"
	// @Param	address	formData	string	false	"收货地址"
	// @Param	phone	formData	string	false	"联系电话"
	// @Param	default	formData	bool	false	"默认地址"

	// 获取用户信息
	user := c.User

	var result resultModels.Result

	//解析 form 表单数据
	var form forms.Address
	//err = c.ParseForm(&form)
	//if err != nil {
	//	c.ResponseErrJson(err)
	//	return
	//}

	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	//根据请求的 id 查找对应地址
	address, err := models.FindAddressById(form.ID)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// ID 不匹配 返回错误信息
	if address.UserId != user.ID {
		result = resultModels.ErrorResult(resultModels.FALL, "记录与用户不匹配，这不是你的地址~")
		c.ResponseJson(result)
		return
	}
	//将表单数据复制到 Address 中
	newAddress := models.Address{}
	err = utils.CopyStruct(form, &newAddress)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	err = models.UpdateAddress(&newAddress)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	if form.Default {
		user.DefaultAddressId = address.ID
		err = models.UpdateUser(user)
		if err != nil {
			c.ResponseErrJson(errors.New("未能成功设置为默认地址"))
			return
		}
	}

	c.ResponseSuccessJson(nil)
	//result = resultModels.SuccessResult(nil)
	//c.ResponseJson(result)
}
