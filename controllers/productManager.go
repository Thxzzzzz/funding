package controllers

import (
	"encoding/json"
	"funding/models"
	"funding/objects"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

// 产品管理相关（商家 审核员）
// 为了方便，这个 Controller 的 URL 地址缩写为 /pm
type ProductMangerController struct {
	VailUserController
}

// @Title 保存产品
// @Description 保存产品
// @Param	form	body	models.Product	true	"产品model"
// @Success 200
// @Failure 400
// @router /save [post]
func (c *ProductMangerController) SaveProduct() {
	// 如果不是卖家，也不是审核员，也不是管理员，那就返回错误
	err := c.VerifySeller()
	if err != nil {
		err = c.VerifyAuditor()
	}
	if err != nil {
		err = c.VerifySuperAdmin()
	}
	if err != nil {
		beego.BeeLogger.Warn(err.Error())
		c.ResponseErrJson(err)
		return
	}

	// 获取传过来的产品信息
	form := models.Product{}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	isUpdate := false
	form.UserId = c.User.ID
	//  如果有 product_id 先查询是否存在 对应产品，存在则说明是要更新而不是新增
	if form.ID > 0 {
		oldProduct, err := models.FindProductById(form.ID)
		// 如果出错返回错误
		if err != nil && err != gorm.ErrRecordNotFound {
			c.ResponseErrJson(err)
			return
			// 如果没找到记录则标记一下，后面将新建产品
		} else if gorm.IsRecordNotFoundError(err) {
			isUpdate = false
		} else {
			// 如果找到了记录，则标记为更新，后面对相应的产品进行更新
			isUpdate = true
		}
		// userId 对不上 也返回错误
		if isUpdate && form.UserId != oldProduct.UserId {
			c.ResponseErrJson(resultError.NewFallFundingErr("这不是你的产品"))
			return
		}
	}

	if isUpdate {
		err = models.UpdateProduct(&form)
	} else {
		err = models.InsertProduct(&form)
	}
	// 发生错误则返回错误
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 成功保存,返回已保存的数据
	c.ResponseSuccessJson(form)
}

// 根据产品 Id 获取产品信息
// @router /productById [get]
func (c *ProductMangerController) GetProductById() {
	// 获取传过来的Id
	id, err := c.GetUint64("id")
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	result, err := models.FindProductById(id)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(result)
}

// 根据产品 Id 获取产品信息
// @router /pkgListByProductId [get]
func (c *ProductMangerController) GetPkgListByProductId() {
	// 获取传过来的Id
	id, err := c.GetUint64("id")
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	result, err := models.FindProductPackagesByProductId(id)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(result)
}
