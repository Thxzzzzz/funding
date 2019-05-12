package models

type Address struct {
	BaseModel
	UserId  uint64 `json:"user_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

// 根据地址的 ID 来获取地址
func FindAddressById(addressId uint64) (*Address, error) {
	var ret Address
	err := db.First(&ret, addressId).Error
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// 根据用户 ID 来获取地址列表
func FindAddressesByUserId(userId uint64) ([]*Address, error) {
	var rets []*Address
	err := db.Find(&rets).Where("user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return rets, nil
}

// 新增地址
func InsertAddress(address *Address) error {
	err := db.Create(address).Error
	return err
}

//删除地址 由于这里是软删除，所以只是把 delete_at 设置了一个值，实际上还存在数据库中,但并不能用 gorm 查到
func DeleteAddressById(id uint64) error {
	err := db.Delete(Address{}, "id = ?", id).Error
	return err
}

//根据 address.ID 来更新其他相应的字段
func UpdateAddress(address *Address) error {
	var rec Address
	err := db.First(&rec, "id = ?", address.ID).Error
	if err != nil {
		return err
	}
	err = db.Model(&rec).Updates(address).Error
	return err
}
