package models

import (
	"fmt"
	"time"
)

//产品
type Product struct {
	BaseModel
	//产品名
	Name string
	//产品类型
	ProductType int
	//当前筹集金额
	CurrentPrice float64
	//目标筹集金额
	TargetPrice float64
	//支持人数
	Backers int
	//截止时间
	EndTime time.Time
	//介绍页详情 Html
	DetailHtml string
	//商品套餐
	ProductPackages []ProductPackage
}

//产品套餐
type ProductPackage struct {
	BaseModel
	//对应产品 Id
	ProductId string
	//套餐描述
	Description string
	//图片链接
	ImageUrl string
	//套餐价格
	Price float64
	//剩余库存
	Stock int64
	//套餐总数
	Total int64
	//运费
	Freight float64
	//发货时间 (众筹成功后多到天内)
	DeliveryDay int64
}

func init() {
	//db.AutoMigrate(&Product{}, &ProductPackage{})
}

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

func GetProductPackage(productId string) (*ProductPackage, error) {
	var result ProductPackage
	err := db.First(&result).Where("Product_id = ?", productId).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
