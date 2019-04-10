package models

import (
	"fmt"
	"funding/objects"
	"time"
)

//产品
type Product struct {
	BaseModel
	//产品名
	Name string `json:"name"`
	//顶部大图
	BigImg string `json:"big_img"`
	//列表小图
	SmallImg string `json:"small_img"`
	//产品类型
	ProductType int `json:"product_type"`
	//当前筹集金额
	CurrentPrice float64 `json:"current_price"`
	//目标筹集金额
	TargetPrice float64 `json:"target_price"`
	//支持人数
	Backers int `json:"backers"`
	//截止时间
	EndTime time.Time `json:"end_time"`
	//介绍页详情 Html
	DetailHtml string `json:"detail_html"`
	//商品套餐
	ProductPackages []ProductPackage `json:"product_packages"`
}

//产品套餐
type ProductPackage struct {
	BaseModel
	//对应产品 Id
	ProductId string `json:"product_id"`
	//套餐描述
	Description string `json:"description"`
	//图片链接
	ImageUrl string `json:"image_url"`
	//套餐价格
	Price float64 `json:"price"`
	//剩余库存
	Stock int64 `json:"stock"`
	//套餐总数
	Total int64 `json:"total"`
	//支持人数
	Backers int `json:"backers"`
	//运费
	Freight float64 `json:"freight"`
	//发货时间 (众筹成功后多少天内)
	DeliveryDay int64 `json:"delivery_day"`
}

func init() {
	//db.AutoMigrate(&Product{}, &ProductPackage{})
}

// 根据 分页 和 产品类型(0 为全部) 获取产品
func GetProductsByPageAndType(page int, pageSize int, productType int) ([]*Product, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, &objects.Error{Msg: "参数错误"}
	}
	var results []*Product
	//分页限制
	pDb := db.Limit(pageSize).Offset((page - 1) * pageSize)
	//类型为 0 时不限制类型
	if productType != 0 {
		pDb = pDb.Where("product_type = ?", productType)
	}
	//倒序查询
	pDb = pDb.Order("id desc")

	err := pDb.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

// 获取全部产品
func GetAllProduct() ([]*Product, error) {
	fmt.Println("Get All Product")
	var results []*Product
	err := db.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func GetProductWithPkg(productId string) (*Product, error) {
	var result Product
	err := db.Preload("ProductPackages").First(&result, productId).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetProductPackages(productId string) (*[]ProductPackage, error) {
	var result []ProductPackage
	err := db.Where("product_id = ?", productId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
