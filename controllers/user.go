package controllers

import (
	"github.com/astaxie/beego"
	"testApi/models"
)

// 用户相关
type UserControllers struct {
	beego.Controller
}

// @Title 根据 id 获取 User
// @Description 根据 Id（数据库表 Id ，不是用户名）来获取对应用户信息
// @Param	id	path	int	true	"数据库 User 表ID"
// @Success 200
// @Failure 400
// @router /id/:id [get]
func (c *UserControllers) GetUserById() {
	// 这里的 Key 要注意带上冒号，否则获取不到对应的参数
	idd, err := c.GetInt64(":id")
	dbResult, err := models.FindUserById(idd)
	var result models.Result
	if err != nil {
		result = models.ErrorResult(models.FALL, err.Error())
	} else {
		result = models.SuccessResult(dbResult)
	}
	c.Data["json"] = result
	c.ServeJSON()
}
