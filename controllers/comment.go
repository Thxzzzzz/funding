package controllers

import (
	"encoding/json"
	"funding/enums"
	"funding/models"
)

// 产品详情获取评论列表是写在 /controller/product.go 里面
// 评论相关 主要是提交评论和回复，需要验证身份
type CommentController struct {
	VailUserController
}

// @Title 保存评论信息
// @Description 保存评论信息
// @Param	form	body	models.CommentsInfo	true	"评论信息"
// @Success 200
// @Failure 400
// @router /saveCommentsInfo [post]
func (c *CommentController) SaveCommentsInfo() {
	// 获取传过来的评论信息
	form := models.CommentsInfo{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	user := c.User
	form.UserId = user.ID
	// 如果是卖家的评论，就去匹配产品，看是否属于该卖家
	if user.RoleId == enums.Role_Seller {
		product, err := models.FindProductById(form.ProductId)
		if err == nil {
			// 匹配上就标记为商家回复
			if product.UserId == user.ID {
				form.IsSeller = true
			}
		}
		err = nil
	}

	// 保存评论
	err = models.InsertCommentsInfo(&form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(form)
}

// @Title 保存回复信息
// @Description 保存回复信息
// @Param	form	body	models.CommentsReply	true	"回复信息"
// @Success 200
// @Failure 400
// @router /saveCommentsReply [post]
func (c *CommentController) SaveCommentsReply() {
	// 获取传过来的评论信息
	form := models.CommentsReply{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	user := c.User
	form.UserId = user.ID
	// 如果是卖家的评论，就去匹配产品，看是否属于该卖家
	if user.RoleId == enums.Role_Seller {
		// 要先查询主评论,获取ProductID
		comment, err := models.FindCommentsInfoById(form.CommentId)
		if err == nil {
			product, err := models.FindProductById(comment.ProductId)
			if err == nil {
				// 匹配上就标记为商家回复
				if product.UserId == user.ID {
					form.IsSeller = true
				}
			}
		}
		err = nil
	}
	// 保存评论
	err = models.InsertCommentsReply(&form)
	if err != nil {
		c.ResponseErrJson(err)
		return
	}
	c.ResponseSuccessJson(form)
}
