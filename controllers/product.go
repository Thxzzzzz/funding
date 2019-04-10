package controllers

import (
	"fmt"
	"funding/models"
	"funding/resultModels"
	"funding/utils"
	"time"
)

// 产品相关
type ProductController struct {
	BaseController
}

// @Title Get Home Page Info
// @Description 获取首页信息
// @Success 200
// @Failure 400
// @router /home
func (c *ProductController) GetHome() {
	var home []resultModels.HomeResult
	var result resultModels.Result
	// 轮播图 前端 type == 0 5个
	//TODO 最新的 5 个产品作为轮播图
	bannerProduct, err := models.GetProductsByPageAndType(1, 5, 0)
	if err != nil {

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
			productContent.CurrentTime = time.Now()
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
	//科技类
	var techType = 1
	techProduct, err := models.GetProductsByPageAndType(1, 7, techType)
	if err != nil {

	} else {
		techResult := resultModels.HomeResult{
			Name:     "科技精选",
			LimitNum: 7,
			Type:     3,
		}
		for _, p := range techProduct {
			var productContent resultModels.ProductContent
			utils.CopyStructJ(&p, &productContent)
			productContent.ID = p.ID
			productContent.CurrentTime = time.Now()
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
//
// @Success 200
// @Failure 400
// @router / [get]
func (c *ProductController) GetProductByPage() {
	// TODO 据页码和其他条件获取产品信息
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
// @Param	id	path	string	true	"商品ID"
// @Success 200
// @Failure 400
// @router /detail/:id [get]
func (c *ProductController) GetProductWithPkg() {
	dbResult, err := models.GetProductWithPkg(c.GetString(":id"))
	var result resultModels.Result
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
	} else {
		result = resultModels.SuccessResult(dbResult)
	}
	c.ResponseJson(result)
}
