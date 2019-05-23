package models

import (
	"funding/forms"
	"funding/objects"
	"github.com/jinzhu/gorm"
	"time"
)

// 主评论表
type CommentsInfo struct {
	BaseModel
	ProductId uint64 `json:"product_id"` // 产品 ID
	UserId    uint64 `json:"user_id"`    // 用户 ID
	IsSeller  bool   `json:"is_seller"`  // 是否是对应卖家
	Content   string `json:"content"`    // 评论内容
}

// 返回给前端的评论信息
type ResultCommentInfo struct {
	ID        uint64               `json:"id" gorm:"primary_key"`
	ProductId uint64               `json:"product_id"` // 产品 ID
	UserId    uint64               `json:"user_id"`    // 用户 ID
	IsSeller  bool                 `json:"is_seller"`  // 是否是对应卖家
	Content   string               `json:"content"`    // 评论内容
	CreatedAt time.Time            `json:"created_at"` // 创建时间
	Username  string               `json:"username"`   // 账号
	Nickname  string               `json:"nickname"`   // 昵称
	IconUrl   string               `json:"icon_url"`   // 用户头像
	Replys    []ResultCommentReply `json:"replys"`     // 回复信息
}

/////////////////////			基本增删改查			/////////////////////

// 根据主评论的 ID 来获取主评论条目
func FindCommentsInfoById(commentsInfoId uint64) (*CommentsInfo, error) {
	var ret CommentsInfo
	err := db.First(&ret, commentsInfoId).Error
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// 根据用户 ID 来获取主评论列表
func FindCommentsInfosByUserId(userId uint64) ([]*CommentsInfo, error) {
	var rets []*CommentsInfo
	err := db.Where("user_id = ?", userId).Find(&rets).Error
	return rets, err
}

// 根据产品 ID 来获取主评论列表
func FindCommentsInfosByProductId(productId uint64) ([]*CommentsInfo, error) {
	var rets []*CommentsInfo
	err := db.Where("product_id = ?", productId).Find(&rets).Error
	return rets, err
}

// 新增主评论
func InsertCommentsInfo(commentsInfo *CommentsInfo) error {
	err := db.Create(commentsInfo).Error
	return err
}

//删除主评论条目 由于这里是软删除，所以只是把 delete_at 设置了一个值，实际上还存在数据库中,但并不能用 gorm 查到
func DeleteCommentsInfoById(id uint64) error {
	err := db.Delete(CommentsInfo{}, "id = ?", id).Error
	return err
}

//根据 commentsInfoID 来更新其他相应的字段
func UpdateCommentsInfo(commentsInfo *CommentsInfo) error {
	var rec CommentsInfo
	err := db.First(&rec, "id = ?", commentsInfo.ID).Error
	if err != nil {
		return err
	}
	err = db.Model(&rec).Updates(commentsInfo).Error
	return err
}

/////////////////////		EMD	基本增删改查			/////////////////////
// 根据产品 ID 获取对应产品的评论
const sqlGetResultCommentInfosByProductId = `
 SELECT
 	c.*,u.username,u.nickname,u.icon_url
 FROM 
 comments_infos c JOIN
 users u ON c.user_id = u.id
 WHERE
 c.deleted_at IS NULL 
 AND u.deleted_at IS NULL 
 AND c.product_id = (?)
 ORDER BY
 id DESC
`

// 获取回复列表
func GetResultCommentInfosByProductId(form *forms.CommentListByProductForm) ([]ResultCommentInfo, error) {
	results := []ResultCommentInfo{}

	//TODO 先不做分页
	//page, pageSize := 1, 10
	//// 如果页码和每页数量大于 0
	//if form.Page > 0 && form.PageSize > 0 {
	//	page = form.Page
	//	pageSize = form.PageSize
	//}

	err := db.Raw(sqlGetResultCommentInfosByProductId, form.ProductId).Scan(&results).Error
	if err != nil {
		return nil, resultError.NewFallFundingErr("查询评论时发生了错误")
	}

	for i, _ := range results {
		replys, err := GetResultCommentReplaysByCommentId(results[i].ID)
		if err != nil && gorm.IsRecordNotFoundError(err) {
			continue
		}
		results[i].Replys = replys
	}
	return results, err
}
