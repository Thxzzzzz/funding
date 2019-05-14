package models

import "funding/objects"

// 评论回复表
type CommentsReplay struct {
	BaseModel
	CommentId uint64 `json:"comment_id"` // 产品 ID
	UserId    uint64 `json:"user_id"`    // 用户 ID
	IsSeller  bool   `json:"is_seller"`  // 是否是对应卖家
	Content   string `json:"content"`    // 评论内容
}

// 返回给前端的回复信息
type ResultCommentReplay struct {
	CommentsInfo
	Username string `json:"username"` // 用户名
	IconUrl  string `json:"icon_url"` // 用户头像
}

/////////////////////			基本增删改查			/////////////////////

// 根据回复的 ID 来获取回复条目
func FindCommentsReplayById(commentsReplayId uint64) (*CommentsReplay, error) {
	var ret CommentsReplay
	err := db.First(&ret, commentsReplayId).Error
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// 根据用户 ID 来获取回复列表
func FindCommentsReplaysByUserId(userId uint64) ([]*CommentsReplay, error) {
	var rets []*CommentsReplay
	err := db.Where("user_id = ?", userId).Find(&rets).Error
	return rets, err
}

// 根据主评论 ID 来获取回复列表
func FindCommentsReplaysByCommentId(commentId uint64) ([]*CommentsReplay, error) {
	var rets []*CommentsReplay
	err := db.Where("comment_id = ?", commentId).Find(&rets).Error
	return rets, err
}

// 新增回复
func InsertCommentsReplay(commentsReplay *CommentsReplay) error {
	err := db.Create(commentsReplay).Error
	return err
}

//删除回复条目 由于这里是软删除，所以只是把 delete_at 设置了一个值，实际上还存在数据库中,但并不能用 gorm 查到
func DeleteCommentsReplayById(id uint64) error {
	err := db.Delete(CommentsReplay{}, "id = ?", id).Error
	return err
}

//根据 commentsReplayID 来更新其他相应的字段
func UpdateCommentsReplay(commentsReplay *CommentsReplay) error {
	var rec CommentsReplay
	err := db.First(&rec, "id = ?", commentsReplay.ID).Error
	if err != nil {
		return err
	}
	err = db.Model(&rec).Updates(commentsReplay).Error
	return err
}

/////////////////////		EMD	基本增删改查			/////////////////////
// 根据产品 ID 获取对应产品的评论
const sqlGetResultCommentReplaysByProductId = `
 SELECT
 	c.*,u.username,u.icon_url
 FROM 
 comments_replays c JOIN
 users u ON c.user_id = u.id
 WHERE
 c.deleted_at IS NULL 
 AND u.deleted_at IS NULL 
 AND c.product_id = (?)
 ORDER BY
 id DESC
`

func GetResultCommentReplaysByProductId(productId uint64) ([]ResultCommentInfo, error) {
	results := []ResultCommentReplay{}
	err := db.Raw(sqlSelectOrderListField, productId).Scan(&results).Error
	if err != nil {
		return nil, resultError.NewFallFundingErr("查询评论时发生了错误")
	}
	for _, comment := range results {

	}
	return results, err
}
