package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"testApi/models"
)

type ProductController struct {
	beego.Controller
}

//@Title Get All Products
//@router /all [get]
func (c *ProductController) GetAll() {
	dbResult, err := models.GetAllProduct()
	var result models.Result
	if err != nil {
		result = models.ErrorResult(models.FALL, err.Error())
	} else {
		result = models.SuccessResult(dbResult)
	}
	fmt.Println(&result)
	c.Data["json"] = result
	c.ServeJSON()
}

//@Title Get Product With Detail
//@router /detail/:id [get]
func (c *ProductController) GetProductWithPkg() {
	dbResult, err := models.GetProductWithPkg(c.GetString("id"))
	var result models.Result
	if err != nil {
		result = models.ErrorResult(models.FALL, err.Error())
	} else {
		result = models.SuccessResult(dbResult)
	}
	c.Data["json"] = result
	c.ServeJSON()
}
