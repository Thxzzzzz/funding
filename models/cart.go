package models

// 购物车相关
type Cart struct {
	BaseModel
	UserId           uint64 `json:"user_id"`            // 用户ID
	ProductPackageId uint64 `json:"product_package_id"` // 套餐ID
	Nums             int    `json:"nums"`               // 购买数量
	Checked          bool   `json:"checked"`            // 是否勾选
}

// 根据购物车的 ID 来获取购物车条目
func FindCartById(cartId uint64) (*Cart, error) {
	var ret Cart
	err := db.First(&ret, cartId).Error
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// 根据用户 ID 来获取购物车列表
func FindCartsByUserId(userId uint64) ([]*Cart, error) {
	var rets []*Cart
	err := db.Find(&rets).Where("user_id = ?", userId).Error
	return rets, err
}

// 新增购物车
func InsertCart(cart *Cart) error {
	err := db.Create(cart).Error
	return err
}

//删除购物车条目 由于这里是软删除，所以只是把 delete_at 设置了一个值，实际上还存在数据库中,但并不能用 gorm 查到
func DeleteCartById(id uint64) error {
	err := db.Delete(Cart{}, "id = ?", id).Error
	return err
}

//根据 cartID 来更新其他相应的字段
func UpdateCart(cart *Cart) error {
	var rec Cart
	err := db.First(&rec, "id = ?", cart.ID).Error
	if err != nil {
		return err
	}
	err = db.Model(&rec).Update(cart).Error
	return err
}

// 根据用户 id 和 套餐 id 来获取对应的购物车条目信息（用于添加到购物差时检查是否已存在）
func FindCartByUserIdAndPkgId(userId uint64, pkdId uint64) (*Cart, error) {
	var ret Cart
	err := db.First(&ret).Where("user_id = ? AND product_package_id = ?", userId, pkdId).Error
	return &ret, err
}
