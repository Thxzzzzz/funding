package managerControllers

import (
	"funding/controllers"
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
	err := c.VerifyAuditor()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
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
	err := c.VerifyAuditor()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
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
