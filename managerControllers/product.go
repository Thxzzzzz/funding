package managerControllers

import (
	"encoding/json"
	"funding/enums"
	"funding/forms"
	"funding/models"
	"funding/objects"
	"github.com/astaxie/beego"
)

type ManagerProductController struct {
	VailManagerController
}

// @Title 获取产品列表，包括未审核的和待提交的也能获取到
// @Description 获取产品列表，包括未审核的 (1：已通过 2：待审核 3: 待提交 4: 未通过 ) 新建的默认应为待审核状态
// @Param	page			query	int		true	"页码"
// @Param	page_size		query	int		true	"每页数量"
// @Param	name			query	string	false	"产品名称"
// @Param	type			query	int		false	"产品类型"
// @Param	funding_status	query	int		false	"众筹状态"
// @Param	sort			query	int		false	"排序方式"
// @Param	price_gt		query	float64	false	"价格大于"
// @Param	price_lt		query	float64	false	"价格小于"
// @Param	verify_status	query	int		false	"审核状态"
// @Success 200
// @Failure 400
// @router /getProductList [get]
func (c *ManagerProductController) GetProductList() {

	// 获取所选的验证状态
	verifyStatus, err := c.GetInt("verify_status")
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}
	form := forms.ProductListForm{}
	// 获取所有 query 数据组成的 map
	values := c.Ctx.Request.URL.Query()
	// 解析到 Struct 中
	err = beego.ParseForm(values, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	form.Page = 1
	form.PageSize = 999

	if err != nil {
		c.ResponseErrJson(err)
	}
	pl, err := models.GetProductList(form, verifyStatus)
	if err != nil {
		c.ResponseErrJson(err)
	}
	c.ResponseSuccessJson(pl)
}

// @Title 更新产品 (包括审核状态)
// @Description  审核状态 ( (1：已通过 2：待审核 3: 待提交 4: 未通过 )  对应 enums.VerifyXXXX 常量)
// @Param	form	body	models.Product	true	"产品model"
// @Success 200
// @Failure 400
// @router /update [post]
func (c *ManagerProductController) UpdateProduct() {
	// 首先要校验权限，是审核人员才能修改审核状态
	// TODO 产品审核状态修改
	// 获取传过来的产品信息
	form := models.Product{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	if form.VerifyStatus == enums.Verify_Success {
		form.VerifyMessage = "审核通过"
	}
	// 更新产品信息
	err = models.UpdateProduct(&form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(nil)
}
