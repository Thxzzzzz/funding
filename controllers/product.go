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
		result = models.ErrorResult(123, err.Error())
	} else {
		result = models.SuccessResult(dbResult)
	}
	fmt.Println(&result)
	c.Data["json"] = result
	c.ServeJSON()
}
