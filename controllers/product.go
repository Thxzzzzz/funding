package controllers

import (
	"fmt"
	"funding/models"
)

// 产品相关
type ProductController struct {
	BaseController
}

// @Title Get All Products
// @Description 获取全部产品信息
// @Success 200
// @Failure 400
// @router /all [get]
func (c *ProductController) GetAll() {
	dbResult, err := models.GetAllProduct()
	var result models.Result
	if err != nil {
		result = models.ErrorResult(models.FALL, err.Error())
	} else {
		result = models.SuccessResult(dbResult)
	}
	fmt.Println(&result)
	c.ResponseJson(result)
}

// @Title 据页码和其他条件获取产品信息
// @Description	据页码和其他条件获取产品信息
//
// @Success 200
// @Failure 400
// @router / [get]
func (c *ProductController) GetProductByPage() {
	// TODO 据页码和其他条件获取产品信息
}

// @Title 根据审核状态获取产品信息
// @Description 审核人员才能调用该接口获取信息 (0:未通过 1：已通过 2：待审核 对应 enums.VerifyXXXX 常量) 新建的默认应为待审核状态
//
// @Success 200
// @Failure 400
// @router /verify [get]
func (c *ProductController) GetVerifyProduct() {
	// 首先要校验权限，是审核人员才能修改审核状态
	// TODO 根据审核状态获取产品信息
}

// @Title 产品审核状态修改
// @Description 审核人员才能修改审核状态 (0:未通过 1：已通过 2：待审核 对应 enums.VerifyXXXX 常量) 新建的默认应为待审核状态
//
// @Success 200
// @Failure 400
// @router /verify/update [post]
func (c *ProductController) VerifyProduct() {
	// 首先要校验权限，是审核人员才能修改审核状态
	// TODO 产品审核状态修改
}

// @Title Get Product With Detail
// @Description 根据 id 获取带有套餐信息的指定产品信息
// @Param	id	path	string	true	"商品ID"
// @Success 200
// @Failure 400
// @router /detail/:id [get]
func (c *ProductController) GetProductWithPkg() {
	dbResult, err := models.GetProductWithPkg(c.GetString(":id"))
	var result models.Result
	if err != nil {
		result = models.ErrorResult(models.FALL, err.Error())
	} else {
		result = models.SuccessResult(dbResult)
	}
	c.ResponseJson(result)
}
