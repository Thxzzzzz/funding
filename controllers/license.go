package controllers

import (
	"encoding/json"
	"funding/enums"
	"funding/models"
	"funding/objects"
	"github.com/jinzhu/gorm"
)

// 营业执照相关
// 放在 /user/license 路由下吧
type LicenseController struct {
	BaseController
}

// @Title 保存产品
// @Description 保存产品
// @Param	form	body	models.License	true	"产品model"
// @Success 200
// @Failure 400
// @router /save [post]
func (c *LicenseController) SaveLicense() {
	// 如果不是卖家 那就返回错误
	user, err := c.CheckAndGetUser()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	//if user.RoleId != enums.Role_Seller {
	//	c.ResponseErrJson(&resultError.UserRoleVerifyError)
	//	return
	//}
	// 获取传过来的产品信息
	form := models.License{}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	isUpdate := false
	// 添加ID
	form.UserId = user.ID
	//  如果有 license_id 先查询是否存在 对应执照，存在则说明是要更新而不是新增
	if form.ID > 0 {
		oldLicense, err := models.FindLicenseById(form.ID)
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

		// 如果是卖家而且 userId 对不上 也返回错误
		if isUpdate && form.UserId != oldLicense.UserId {
			c.ResponseErrJson(resultError.NewFallFundingErr("这不是你的执照"))
			return
		}
	}

	// 卖家不能改成除了待审核/待提交以外的状态，如果这以外的参数有则去掉
	if form.VerifyStatus != enums.Verify_Wait && form.VerifyStatus != enums.Verify_UnSubmit {
		form.VerifyStatus = 0
	}

	if isUpdate {
		err = models.UpdateLicense(&form)
	} else {
		err = models.InsertLicense(&form)
	}
	// 发生错误则返回错误
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	// 成功保存,返回已保存的数据
	c.ResponseSuccessJson(form)
}
