package managerControllers

import (
	"encoding/json"
	"funding/controllers"
	"funding/enums"
	"funding/models"
	"funding/objects"
	"github.com/jinzhu/gorm"
)

// 执照信息相关的 API
// BASE_API 应该为 /manager/ 但是自动生成的是错的，不知道怎么改,所以这个端口 SWAGGER 不能正常测试
type LicenseController struct {
	controllers.VailUserController
}

// @Title 获取全部执照信息
// @Description
// @Success 200
// @Failure 400
// @router /all [get]
func (c *LicenseController) GetAllLicense() {
	// 验证是否是审核员
	//err := c.VerifyAuditor()
	//if err != nil {
	//	c.ResponseErrJson(err)
	//	return
	//}
	// 查询所有执照信息
	results, err := models.GetAllLicense()
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(results)
}

// @Title 获取对应状态的执照信息
// @Description
// @Param	verify_status	query	int	true	"验证状态"
// @Success 200
// @Failure 400
// @router /getByVerifyStatus [get]
func (c *LicenseController) GetByVerifyStatus() {
	// 验证是否是审核员
	//err := c.VerifyAuditor()
	//if err != nil {
	//	c.ResponseErrJson(err)
	//	return
	//}
	verifyStatus, err := c.GetInt("verify_status")
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}
	// 查询对应状态的执照信息
	results, err := models.GetLicenseByVerifyStatus(verifyStatus)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(results)
}

// @Title 更新执照 (包括审核状态)
// @Description  审核状态 ( (1：已通过 2：待审核 3: 待提交 4: 未通过 )  对应 enums.VerifyXXXX 常量)
// @Param	form	body	models.License	true	"执照model"
// @Success 200
// @Failure 400
// @router /update [post]
func (c *LicenseController) UpdateLicense() {
	// 首先要校验权限，是审核人员才能修改审核状态
	// TODO 产品审核状态修改
	// 获取传过来的产品信息
	form := models.License{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}

	// 通过验证，将对应用户改为商家
	if form.VerifyStatus == enums.Verify_Success {
		form.VerifyMessage = "审核通过"
		user, err := models.FindUserById(form.ID)
		if err != nil {
			c.ResponseErrJson(err)
			return
		}
		// 通过验证，将对应用户改为商家
		if user.RoleId != enums.Role_Seller {
			user.RoleId = enums.Role_Seller
			err = models.UpdateUser(user)
			if err != nil {
				c.ResponseErrJson(err)
				return
			}
		}
	}
	// 更新产品信息
	err = models.UpdateLicense(&form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(nil)
}
