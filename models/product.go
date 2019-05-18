package models

import (
	"errors"
	"funding/enums"
	"funding/forms"
	"funding/resultModels"
	"strings"
	"time"
)

//产品
type Product struct {
	BaseModel
	Name            string              `json:"name"`                           //产品名
	BigImg          string              `json:"big_img"`                        //顶部大图
	SmallImg        string              `json:"small_img"`                      //列表小图
	UserId          uint64              `json:"user_id"`                        //发布者ID
	ProductType     int                 `json:"product_type"`                   //产品类型
	CurrentPrice    float64             `json:"current_price"`                  //当前筹集金额
	TargetPrice     float64             `json:"target_price"`                   //目标筹集金额
	VerifyStatus    int                 `json:"verify_status" gorm:"default:3"` //审核状态(1：已通过 2：待审核 3:待提交（默认） 4:未通过 )
	VerifyMessage   string              `json:"verify_message"`                 //审核消息（审核失败的原因）
	FundingStatus   enums.FundingStatus `json:"funding_status"  gorm:"-"`       //众筹状态 TODO:(现在这个字段数据库里还没有)
	Backers         int                 `json:"backers"`                        //支持人数
	EndTime         time.Time           `json:"end_time"`                       //截止时间
	DetailHtml      string              `json:"detail_html"`                    //介绍页详情 Html
	ProductPackages []ProductPackage    `json:"product_packages"`               //商品套餐
}

// 产品类型
type ProductType struct {
	ID   int    `json:"id" gorm:"primary_key"` //类型 id
	Name string `json:"name"`                  //类型名称
}

func init() {
	//db.AutoMigrate(&Product{}, &ProductPackage{})
}

// 计算众筹状态
func CalcFundingStatus(currentTime time.Time, endTime time.Time,
	current_price float64, target_price float64) enums.FundingStatus {
	if currentTime.Before(endTime) {
		return enums.FundingStatus_Ing
	}
	if current_price >= target_price {
		return enums.FundingStatus_Success
	}
	return enums.FundingStatus_Fail
}

/////////////////////			基本增删改查			/////////////////////

// 根据产品的 ID 来获取产品条目
func FindProductById(productId uint64) (*Product, error) {
	var ret Product
	err := db.First(&ret, productId).Error
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// 根据用户 ID 来获取产品列表
func FindProductsByUserId(userId uint64) ([]*Product, error) {
	var result []*Product
	err := db.Where("user_id = ?", userId).Order("id DESC").Find(&result).Error
	if err != nil {
		return nil, err
	}
	timeNow := time.Now()
	for i := range result {
		result[i].FundingStatus = CalcFundingStatus(timeNow, result[i].EndTime,
			result[i].CurrentPrice, result[i].TargetPrice)
	}
	return result, err
}

//func FinProductByVerifyStatus()

// 新增产品
func InsertProduct(product *Product) error {
	err := db.Create(product).Error
	return err
}

//删除产品条目 由于这里是软删除，所以只是把 delete_at 设置了一个值，实际上还存在数据库中,但并不能用 gorm 查到
func DeleteProductById(id uint64) error {
	err := db.Delete(Product{}, "id = ?", id).Error
	return err
}

//根据 productID 来更新其他相应的字段
func UpdateProduct(product *Product) error {
	var rec Product
	err := db.First(&rec, "id = ?", product.ID).Error
	if err != nil {
		return err
	}
	err = db.Model(&rec).Updates(product).Error
	return err
}

// 获取全部产品
func GetAllProduct() ([]*Product, error) {
	var results []*Product
	err := db.Find(&results).Error
	return results, err
}

/////////////////////		EMD	基本增删改查			/////////////////////

// 获取产品类型列表
func GetProductTypeList() ([]ProductType, error) {
	results := []ProductType{}
	err := db.Find(&results).Error
	return results, err
}

// 根据 分页 和 产品类型(0 为全部) 获取产品
func GetProductsByPageAndType(page int, pageSize int, productType int) ([]*Product, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, errors.New("参数错误")
	}
	var results []*Product
	// 分页限制
	pDb := db.Limit(pageSize).Offset((page - 1) * pageSize)
	// 类型为 0 时不限制类型
	if productType != 0 {
		pDb = pDb.Where("product_type = ?", productType)
	}
	// 只查询通过验证的
	pDb = pDb.Where("verify_status = ?", enums.Verify_Success)
	// 倒序查询
	pDb = pDb.Order("created_at DESC,end_time DESC")
	// 查询所有符合条件的数据返回到 results 里面
	err := pDb.Find(&results).Error
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()
	for i := range results {
		results[i].FundingStatus = CalcFundingStatus(timeNow, results[i].EndTime,
			results[i].CurrentPrice, results[i].TargetPrice)
	}
	return results, nil
}

func GetProductWithPkg(productId uint64) (*Product, error) {
	var result Product
	err := db.Preload("ProductPackages").First(&result, productId).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// 根据页码等信息获取产品列表
func GetProductList(form forms.ProductListForm, verifyStatus int) (*resultModels.ProductList, error) {
	result := resultModels.ProductList{}
	var list []*resultModels.ProductContent
	cDb := db
	// 如果传入类型不为 0 （传入了 Type)，则根据 Type 查询
	if form.Type != 0 {
		cDb = cDb.Where("product_type = ?", form.Type)
	}
	form.Name = strings.TrimSpace(form.Name)
	// 传入的名称不为空，则根据名称查询
	if form.Name != "" {
		cDb = cDb.Where("name LIKE ?", "%"+form.Name+"%")
	}

	if form.Type != 0 {
		cDb = cDb.Where("name LIKE ?", "%"+form.Name+"%")
	}
	page, pageSize := 1, 10
	// 如果页码和每页数量大于 0
	if form.Page > 0 && form.PageSize > 0 {
		page = form.Page
		pageSize = form.PageSize
	}
	if verifyStatus != 0 {
		// 只查询指定验证状态
		cDb = cDb.Where("verify_status = ?", verifyStatus)
	}
	// 只查询指定字段
	cDb = cDb.Select(resultModels.ProductContentField)

	// 未软删除的行
	cDb = cDb.Where("deleted_at IS NULL").Table("products")
	// 排序
	cDb = cDb.Order("created_at DESC,end_time DESC")
	// 统计总数
	err := cDb.Count(&result.Total).Error
	if err != nil {
		return nil, err
	}
	// 调整 Offset(偏移量，控制页数),和 Limit (数量限制，控制每页数量）
	cDb = cDb.Offset((page - 1) * pageSize).Limit(pageSize)
	err = cDb.Scan(&list).Error

	timeNow := time.Now()
	for i := range list {
		list[i].FundingStatus = CalcFundingStatus(timeNow, list[i].EndTime,
			list[i].CurrentPrice, list[i].TargetPrice)
	}
	result.Page = form.Page
	result.ProductContents = list
	return &result, err
}

const sqlGetCheckoutPkgInfo = `
SELECT
	p.id AS product_id,pkg.id AS product_package_id,
	p.name AS product_name, pkg.price AS price,
	p.user_id AS seller_id ,u.nickname AS seller_nickname,
	pkg.image_url,pkg.description,
	pkg.stock,p.end_time
FROM
	products p LEFT JOIN
	product_packages pkg ON p.id = pkg.product_id
	JOIN users u ON p.user_id = u.id
WHERE
	p.deleted_at IS NULL 
	AND pkg.deleted_at IS NULL
	AND p.verify_status = 1
	AND pkg.id = (?) 
`

// 从套餐 ID 获取所需的结算信息
func GetCheckoutPkgInfoFromPkgId(pkgId uint64) (*resultModels.CheckoutPkgInfo, error) {
	var result resultModels.CheckoutPkgInfo
	// 根据 SQL 字符串拼接查询订单相关信息列表
	err := db.Raw(sqlGetCheckoutPkgInfo, pkgId).Scan(&result).Error
	return &result, err
}

const sqlCountProduct = `
SUM(current_price) AS support_price_count,
MAX(current_price) AS max_support_price,
SUM(backers) AS backers_count,
MAX(backers) AS max_backers
`

// 统计产品信息
func GetAllProductCountInfo() (resultModels.ProductCountInfo, error) {
	countInfo := resultModels.ProductCountInfo{}
	err := db.Table("products").Select(sqlCountProduct).Scan(&countInfo).Error
	return countInfo, err
}

// 获取产品的截止日期,这个可以用作购物车的失效处理,或者在获取购物车列表的时候就处理？
//func GetEndTimeListInProductId(productIds []uint64) ([]time.Time, error) {
//	results := []time.Time{}
//	err := db.Select("end_time").Table("products").Where("id IN (?)", productIds).Error
//	return results, err
//}
