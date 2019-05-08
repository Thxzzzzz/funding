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
const sqlGetCartItemsByUserId = `
SELECT
	c.id,c.user_id,p.user_id AS seller_id,c.product_package_id,c.nums,c.checked,pkg.product_id,
	p.name AS product_name,pkg.price,pkg.stock,pkg.image_url,pkg.description,p.end_time
FROM
	carts c
JOIN
	product_packages pkg ON c.product_package_id = pkg.id
JOIN
	products p ON pkg.product_id = p.id
WHERE
	c.deleted_at IS NULL  AND
	p.deleted_at IS NULL  AND
	pkg.deleted_at IS NULL AND
	c.user_id = (?) 
`

// 根据 UserId 和 Product_package_id  查询前端所对应的 CartItem SQL 语句
const sqlGetCartItemByUserIdAndPkgId = sqlGetCartItemsByUserId + `AND	c.product_package_id = (?)`

const sqlOrderByCartUpdateDesc = `ORDER BY c.updated_at DESC`

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

	// 由于 Checked 是 bool 型其默认值是 0 ，在直接传入 struct 更新时会被 Gorm 忽略掉，所以这里要手动写更新的 Map
	err = db.Model(&rec).Updates(map[string]interface{}{"checked": cart.Checked, "nums": cart.Nums}).Error

	//err = db.Model(&rec).Update(cart).Error
	return err
}

// 根据用户 id 和 套餐 id 来获取对应的购物车条目信息（用于添加到购物差时检查是否已存在）
func FindCartByUserIdAndPkgId(userId uint64, pkgId uint64) (*Cart, error) {
	var ret Cart
	err := db.Where("user_id = ? AND product_package_id = ?", userId, pkgId).First(&ret).Error
	return &ret, err
}

// 根据用户 id 和 套餐 id 来更新对应的购物车条目信息
func UpdateCartByUserIdAndPkgId(cart *Cart) error {
	ret, err := FindCartByUserIdAndPkgId(cart.UserId, cart.ProductPackageId)
	if err != nil {
		return err
	}
	ret.Checked = cart.Checked
	ret.Nums = cart.Nums
	err = UpdateCart(ret)
	return err
}

// 根据用户 id 和 套餐 id 来删除对应购物车信息
func DeleteCartByUserIdAndPkgId(userId uint64, pkgId uint64) error {
	err := db.Table("carts").Where("deleted_at IS NULL AND user_id = ? AND product_package_id = ?", userId, pkgId).Delete(Cart{}).Error
	return err
}

// 返回购物车列表
func GetCartItems(userId uint64) ([]resultModels.CartItem, error) {
	var results []resultModels.CartItem
	// 执行 SQL 语句，并将结果映射到 results 中
	err := db.Raw(sqlGetCartItemsByUserId+sqlOrderByCartUpdateDesc, userId).Scan(&results).Error
	return results, err
}

// 返回购物车列表项目
func GetCartItemByUserIdAndPkgId(userId uint64, pkgId uint64) (resultModels.CartItem, error) {
	var result resultModels.CartItem
	// 执行 SQL 语句，并将结果映射到 result 中
	err := db.Raw(sqlGetCartItemByUserIdAndPkgId+sqlOrderByCartUpdateDesc, userId, pkgId).Scan(&result).Error
	return result, err
}

// 全选/全不选
func UpdateAllCheckedStatus(userId uint64, checked bool) error {
	err := db.Table("carts").Where("deleted_at IS NULL AND user_id = ? ", userId).Updates(map[string]interface{}{"checked": checked}).Error
	return err
}

// 删除所有选中项
func DeleteAllCheckedCarts(userId uint64) error {
	err := db.Table("carts").Where("deleted_at IS NULL AND checked = true AND user_id = ? ", userId).Delete(Cart{}).Error
	return err
}
