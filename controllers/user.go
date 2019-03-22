package controllers

import (
	"github.com/astaxie/beego"
	"testApi/models"
)

type UserControllers struct {
	beego.Controller
}

//TODO: 这里不知道为什么，获取不到参数
//@Title 根据 id 获取 User
//@router  /id/:id [get]
func (c *UserControllers) GetUserById() {
	id, err := c.GetInt("id")
	dbResult, err := models.FindUserById(id)
	var result models.Result
	if err != nil {
		result = models.ErrorResult(models.FALL, err.Error())
	} else {
		result = models.SuccessResult(dbResult)
	}
	c.Data["json"] = result
	c.ServeJSON()
}
