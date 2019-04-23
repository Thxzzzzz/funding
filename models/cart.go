package models

import "funding/resultModels"

// 购物车相关
type Cart struct {
	BaseModel
	UserId           uint64 `json:"user_id"`            // 用户ID
	ProductPackageId uint64 `json:"product_package_id"` // 套餐ID
	Nums             int    `json:"nums"`               // 购买数量
	Checked          bool   `json:"checked"`            // 是否勾选
}

// 根据 UserId 查询前端所需要的 CartItem 列表 SQL 语句
const sqlGetCartItemsByUserId = `SELECT
c.*,pkg.product_id,p.name,pkg.image_url,pkg.description
FROM
	carts c
JOIN
	product_packages pkg ON c.product_package_id = pkg.id
JOIN
	products p ON pkg.product_id = p.id
WHERE
	c.user_id = ? 
`

// 根据 UserId 和 Product_package_id  查询前端所对应的 CartItem SQL 语句
const sqlGetCartItemByUserIdAndPkgId = sqlGetCartItemsByUserId + `AND	c.product_package_id = ?`

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
	err := db.Where("user_id = ? AND product_package_id = ?", userId, pkdId).First(&ret).Error
	return &ret, err
}

// 返回购物车列表
func GetCartItems(userId uint64) ([]resultModels.CartItem, error) {
	var results []resultModels.CartItem
	// 执行 SQL 语句，并将结果映射到 results 中
	err := db.Raw(sqlGetCartItemsByUserId, userId).Scan(&results).Error
	return results, err
}

// 返回购物车列表项目
func GetCartItemByUserIdAndPkgId(userId uint64, pkgId uint64) (resultModels.CartItem, error) {
	var result resultModels.CartItem
	// 执行 SQL 语句，并将结果映射到 result 中
	err := db.Raw(sqlGetCartItemByUserIdAndPkgId, userId, pkgId).Scan(&result).Error
	return result, err
}
