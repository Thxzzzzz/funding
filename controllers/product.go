package controllers

import (
	"fmt"
	"funding/forms"
	"funding/models"
	"funding/objects"
	"funding/resultModels"
	"funding/utils"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

// 产品相关
type ProductController struct {
	BaseController
}

// @Title 获取产品类型列表
// @Description 获取产品类型列表
// @Success 200 {object} []models.ProductType
// @Failure 400
// @router /typeList [get]
func (c *ProductController) GetProductTypeList() {
	results, err := models.GetProductTypeList()
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(results)
}

// @Title Get Home Page Info
// @Description 获取首页信息
// @Success 200
// @Failure 400
// @router /home [get]
func (c *ProductController) GetHome() {
	var home []resultModels.HomeResult
	var result resultModels.Result
	// 轮播图 前端 type == 0 5个
	//TODO 最新的 5 个产品作为轮播图
	bannerProduct, err := models.GetProductsByPageAndType(1, 5, 0)
	if err != nil {
		c.ResponseErrJson(err)
		return
	} else {
		homeBanner := resultModels.HomeResult{
			Name:     "轮播图",
			LimitNum: 5,
			Type:     0,
		}
		for _, p := range bannerProduct {
			var productContent resultModels.ProductContent
			utils.CopyStructJ(&p, &productContent)
			productContent.ID = p.ID
			homeBanner.ProductContents = append(homeBanner.ProductContents, productContent)
		}
		home = append(home, homeBanner)
	}

	// 活动板块  前端 type == 1
	//TODO 这个感觉不需要。。

	// 热门商品  前端 type == 2 2个？
	//TODO 众筹中的产品里面筹集金额最高的产品

	// XXX精选 前端 type == 3 7个
	//TODO 几大类别的热门
	productTypes, err := models.GetProductTypeList()
	if err != nil {
		c.ResponseErrJson(resultError.NewFallFundingErr("获取产品类型失败"))
		return
	}
	for _, productType := range productTypes {
		//var techType = 1
		products, err := models.GetProductsByPageAndType(1, 7, productType.ID)
		if err != nil && err != gorm.ErrRecordNotFound {
			c.ResponseErrJson(err)
			return
		}
		// 这个类别没有记录就跳到下一个
		if err == gorm.ErrRecordNotFound {
			continue
		}
		// 小于 6 个 就跳过
		if len(products) < 6 {
			continue
		}
		techResult := resultModels.HomeResult{
			Name:     productType.Name + "精选",
			LimitNum: 7,
			Type:     3,
		}

		for _, p := range products {
			var productContent resultModels.ProductContent
			utils.CopyStructJ(&p, &productContent)
			productContent.ID = p.ID
			techResult.ProductContents = append(techResult.ProductContents, productContent)
		}
		home = append(home, techResult)

	}

	result = resultModels.SuccessResult(home)
	c.ResponseJson(result)
}

// @Title Get All Products
// @Description 获取全部产品信息
// @Success 200
// @Failure 400
// @router /all [get]
func (c *ProductController) GetAll() {
	dbResult, err := models.GetAllProduct()
	var result resultModels.Result
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
	} else {
		result = resultModels.SuccessResult(dbResult)
	}
	fmt.Println(&result)
	c.ResponseJson(result)
}

// @Title 据页码和其他条件获取产品信息
// @Description	据页码和其他条件获取产品信息
// @Param	page			query	int		true	"页码"
// @Param	page_size		query	int		true	"每页数量"
// @Param	type			query	int		false	"产品类型"
// @Param	status			query	int		false	"众筹状态"
// @Param	sort			query	int		false	"排序方式"
// @Param	price_gt		query	float64	false	"价格大于"
// @Param	price_lt		query	float64	false	"价格小于"
// @Success 200
// @Failure 400
// @router /productList [get]
func (c *ProductController) GetProductByPage() {
	// TODO 据页码和其他条件获取产品信息
	form := forms.ProductListForm{}
	// 获取所有 query 数据组成的 map
	values := c.Ctx.Request.URL.Query()
	// 解析到 Struct 中
	err := beego.ParseForm(values, &form)
	if err != nil {
		c.ResponseErrJson(err)
	}
	pl, err := models.GetProductList(form)
	if err != nil {
		c.ResponseErrJson(err)
	}
	c.ResponseSuccessJson(pl)
}

// @Title 根据审核状态获取产品信息
// @Description 审核人员才能调用该接口获取信息 (0:未通过 1：已通过 2：待审核 对应 enums.VerifyXXXX 常量) 新建的默认应为待审核状态
//
// @Success 200
// @Failure 400
// @router /verify [get]
func (c *ProductController) GetVerifyProduct() {
	// 首先要校验权限，是审核人员才能修改审核状态
	// TODO 根据审核状态获取产品信息
}

// @Title 产品审核状态修改
// @Description 审核人员才能修改审核状态 (0:未通过 1：已通过 2：待审核 对应 enums.VerifyXXXX 常量) 新建的默认应为待审核状态
//
// @Success 200
// @Failure 400
// @router /verify/update [post]
func (c *ProductController) VerifyProduct() {
	// 首先要校验权限，是审核人员才能修改审核状态
	// TODO 产品审核状态修改
}

// @Title Get Product With Detail
// @Description 根据 id 获取带有套餐信息的指定产品信息
// @Param	id	query	int	true	"商品ID"
// @Success 200	{object} models.Product
// @Failure 400
// @router /detail [get]
func (c *ProductController) GetProductWithPkg() {
	id, err := c.GetUint64("id")
	var result resultModels.Result
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
	}
	dbResult, err := models.GetProductWithPkg(id)
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
	} else {
		result = resultModels.SuccessResult(dbResult)
	}
	c.ResponseJson(result)
}

// @Title 获取结算所需的套餐信息 （给“立即支持”这个功能用）
// @Description 获取结算所需的套餐信息 （给“立即支持”这个功能用）
// @Param	product_package_id	query	int	true	"套餐ID"
// @Success	200 {object} resultModels.CheckoutPkgInfo
// @Failure 400
// @router /checkoutPkgInfo [get]
func (c *ProductController) GetCheckoutPkgInfo() {
	// 从请求参数中获取 product_package_id 的值
	pkgId, err := c.GetUint64("product_package_id")
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}
	// 到数据库中查询
	result, err := models.GetCheckoutPkgInfoFromPkgId(pkgId)

	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(result)
}

// @Title 根据商品 ID 获取商家信息
// @Description 根据商品 ID 获取商家信息
// @Param	product_id	query	int	true	"套餐ID"
// @Success	200
// @Failure 400
// @router /getSellerByProductId [get]
func (c *ProductController) GetSellerByProductId() {
	productId, err := c.GetUint64("product_id")
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}
	// 查询对应订单
	product, err := models.FindProductById(productId)
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}
	// 查询对应卖家
	user, err := models.FindUserById(product.UserId)
	if err != nil {
		c.ResponseErrJson(&resultError.FormParamErr)
		return
	}
	fmt.Println(user)
	//TODO 查询对应的License信息

}

// @Title 获取统计信息
// @Description 获取统计信息
// @Success	200
// @Failure 400
// @router /getProductCountInfo [get]
func (c *ProductController) GetProductCountInfo() {

	//TODO 获取统计信息
}
