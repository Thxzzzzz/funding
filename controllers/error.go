package controllers

import (
	"funding/resultModels"
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	result := resultModels.ErrorResult(404, "Api Not Found")
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *ErrorController) Error501() {
	result := resultModels.ErrorResult(501, "Server Error")
	c.Data["json"] = result
	c.ServeJSON()
}
