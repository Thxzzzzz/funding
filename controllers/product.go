package controllers

import (
	"fmt"
	"testApi/models"
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
	ss := c.GetSession("userId")
	fmt.Println(ss)
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
