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
func InsertAddress(address Address) error {

	return nil
}
