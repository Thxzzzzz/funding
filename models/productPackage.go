package models

//产品套餐
type ProductPackage struct {
	BaseModel
	ProductId   string  `json:"product_id"`   //对应产品 Id
	Description string  `json:"description"`  //套餐描述
	ImageUrl    string  `json:"image_url"`    //图片链接
	Price       float64 `json:"price"`        //套餐价格
	Stock       int64   `json:"stock"`        //剩余库存
	Total       int64   `json:"total"`        //套餐总数
	Backers     int     `json:"backers"`      //支持人数
	Freight     float64 `json:"freight"`      //运费
	DeliveryDay int64   `json:"delivery_day"` //发货时间 (众筹成功后多少天内)
}

/////////////////////			基本增删改查			/////////////////////

// 根据产品套餐的 ID 来获取产品套餐条目
func FindProductPackageById(productPackageId uint64) (*ProductPackage, error) {
	var ret ProductPackage
	err := db.First(&ret, productPackageId).Error
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// 根据用户 ID 来获取产品套餐列表
func FindProductPackagesByUserId(userId uint64) ([]*ProductPackage, error) {
	var rets []*ProductPackage
	err := db.Find(&rets).Where("user_id = ?", userId).Error
	return rets, err
}

// 新增产品套餐
func InsertProductPackage(productPackage *ProductPackage) error {
	err := db.Create(productPackage).Error
	return err
}

//删除产品套餐条目 由于这里是软删除，所以只是把 delete_at 设置了一个值，实际上还存在数据库中,但并不能用 gorm 查到
func DeleteProductPackageById(id uint64) error {
	err := db.Delete(ProductPackage{}, "id = ?", id).Error
	return err
}

//根据 productPackageID 来更新其他相应的字段
func UpdateProductPackage(productPackage *ProductPackage) error {
	var rec ProductPackage
	err := db.First(&rec, "id = ?", productPackage.ID).Error
	if err != nil {
		return err
	}
	err = db.Model(&rec).Updates(productPackage).Error
	return err
}

/////////////////////		EMD	基本增删改查			/////////////////////
