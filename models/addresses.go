package models

type Address struct {
	BaseModel
	UserId  uint64 `json:"user_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func GetAddressById(addressId uint64) (*Address, error) {
	var ret Address
	err := db.First(&ret, addressId).Error
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func GetAddressesByUserId(userId uint64) ([]*Address, error) {
	var rets []*Address
	err := db.Find(&rets).Where("user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return rets, nil
}
