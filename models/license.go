package models

// 执照信息 对应 licenses 表
type License struct {
	BaseModel
	UserId          uint64 `json:"user_id"`           // 用户ID
	CompanyName     string `json:"company_name"`      // 公司名
	Description     string `json:"description"`       // 公司描述
	Address         string `json:"address"`           // 联系地址
	Phone           string `json:"phone"`             // 联系电话
	LicenseImageUrl string `json:"license_image_url"` // 营业执照照片地址
	VerifyStatus    int    `json:"verify_status"`     // 审核状态  1：已通过 2：待审核 3: 待提交 4: 未通过
	VerifyMessage   string `json:"verify_message"`    // 审核消息（审核不通过时显示）
}

/////////////////////			基本增删改查			/////////////////////

// 根据执照信息的 ID 来获取执照信息条目
func FindLicenseById(licenseId uint64) (*License, error) {
	var ret License
	err := db.First(&ret, licenseId).Error
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// 根据用户 ID 来获取
func FindLicensesByUserId(userId uint64) (License, error) {
	var ret License
	err := db.Where("user_id = ?", userId).First(&ret).Error
	return ret, err
}

// 新增执照信息
func InsertLicense(license *License) error {
	err := db.Create(license).Error
	return err
}

//删除执照信息条目 由于这里是软删除，所以只是把 delete_at 设置了一个值，实际上还存在数据库中,但并不能用 gorm 查到
func DeleteLicenseById(id uint64) error {
	err := db.Delete(License{}, "id = ?", id).Error
	return err
}

//根据 licenseID 来更新其他相应的字段
func UpdateLicense(license *License) error {
	var rec License
	err := db.First(&rec, "id = ?", license.ID).Error
	if err != nil {
		return err
	}
	err = db.Model(&rec).Updates(license).Error
	return err
}

// 获取全部执照信息
func GetAllLicense() ([]*License, error) {
	var results []*License
	err := db.Find(&results).Error
	return results, err
}

/////////////////////		EMD	基本增删改查			/////////////////////

// 根据审核状态获取执照列表
func GetLicenseByVerifyStatus(verifyStatus int) ([]*License, error) {
	var results []*License
	err := db.Where("verify_status = ?", verifyStatus).Find(&results).Error
	return results, err
}
