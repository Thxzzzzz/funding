package controllers

import (
	"encoding/json"
	"errors"
	"funding/enums"
	"funding/forms"
	"funding/models"
	"funding/objects"
	"funding/utils"
	"github.com/jinzhu/gorm"
)

// 用户购物车相关
type CartController struct {
	VailUserController
}

/////////						 Carts 购物车相关   									///////////

// @Title 购物车列表
// @Description	获取购物车列表
// @Success 200 {object} []models.Cart
// @Failure 400
// @router /cartList [get]
func (c *CartController) CartList() {
	// 获取用户信息
	user := c.User
	rec, err := models.GetCartItems(user.ID)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	c.ResponseSuccessJson(rec)
}

// @Title 添加购物车
// @Description 添加购物车
// @Param	cartForm	body	forms.CartForm	true	"购物车信息"
// @Success 200
// @Failure 400
// @router /addCart [post]
func (c *CartController) AddCart() {
	// 获取用户信息
	user := c.User
	// 不是买家就不加购物车了吧
	if user.RoleId != enums.Role_Buyer {
		c.ResponseErrJson(resultError.NewFallFundingErr("这不是买家"))
		return
	}
	//解析 form 表单数据
	var form forms.CartForm
	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// 2. 根据 userId 和 product_package_id 来检查是否存在对应的购物车记录
	rec, err := models.FindCartByUserIdAndPkgId(user.ID, form.ProductPackageId)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		c.ResponseErrJson(err)
		return
	}

	// 3.1 存在对应记录则对其数量进行增加
	if rec.ID != 0 {
		rec.Nums += form.Nums
		err = models.UpdateCart(rec)
	} else { // 3.2 不存在则插入新的购物车记录
		// 将表单数据复制到对应 model 中
		cart := models.Cart{}
		err = utils.CopyStruct(form, &cart)
		// 把对应的 UserId 加入到数据中
		cart.UserId = user.ID
		if err != nil {
			c.ResponseErrJson(err)
			return
		}
		err = models.InsertCart(&cart)
		if err != nil {
			c.ResponseErrJson(err)
			return
		}
	}
	c.ResponseSuccessJson(nil)
}

// @Title 删除当前用户指定 product_package_id 的购物车记录
// @Description 删除当前用户指定 product_package_id 的购物车记录
// @Param	delCartForm	body	forms.DelCartForm	true	"购物车信息"
// @Success 200
// @Failure 400
// @router /cartDel [post]
func (c *CartController) CartDel() {
	// 获取用户信息
	user := c.User
	//解析 form 表单数据
	var form forms.DelCartForm
	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	err = models.DeleteCartByUserIdAndPkgId(user.ID, form.ProductPackageId)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(nil)
}

// @Title 编辑购物车
// @Description 编辑购物车
// @Param	cartForm	body	forms.CartForm	true	"购物车信息"
// @Success 200
// @Failure 400
// @router /cartEdit [post]
func (c *CartController) CartEdit() {
	// 获取用户信息
	user := c.User
	//解析 form 表单数据
	var form forms.CartForm
	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// 2. 根据 userId 和 product_package_id 来检查是否存在对应的购物车记录
	rec, err := models.FindCartByUserIdAndPkgId(user.ID, form.ProductPackageId)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			c.ResponseErrJson(errors.New("购物车记录不存在"))
			return
		}
		c.ResponseErrJson(err)
		return
	}

	// 3.1 存在对应记录则对其数量进行增加
	rec.Nums = form.Nums
	rec.Checked = form.Checked
	err = models.UpdateCart(rec)
	c.ResponseSuccessJson(nil)
}

// @Title 全选/全不选
// @Description 全选/全不选
// @Param	cartForm	body	forms.CheckAllForm	true	"购物车信息"
// @Success 200
// @Failure 400
// @router /editCheckAll [post]
func (c *CartController) EditCheckAll() {
	// 获取用户信息
	user := c.User
	//解析 form 表单数据
	var form forms.CheckAllForm
	//这里由于 前端的 Axios 默认请求为 json 格式，所以先改为解析Json
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	err = models.UpdateAllCheckedStatus(user.ID, form.Checked)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(nil)
}

// @Title 删除选中
// @Description 删除选中
// @Success 200
// @Failure 400
// @router /delCartChecked [post]
func (c *CartController) DelCartChecked() {
	// 获取用户信息
	user := c.User
	err := models.DeleteAllCheckedCarts(user.ID)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(nil)
}
