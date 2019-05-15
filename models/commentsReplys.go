package models

import "time"

// 评论回复表
type CommentsReply struct {
	BaseModel
	CommentId uint64 `json:"comment_id"` // 产品 ID
	UserId    uint64 `json:"user_id"`    // 用户 ID
	IsSeller  bool   `json:"is_seller"`  // 是否是对应卖家
	Content   string `json:"content"`    // 评论内容
}

// 返回给前端的回复信息
type ResultCommentReply struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	CommentId uint64    `json:"comment_id"` // 产品 ID
	UserId    uint64    `json:"user_id"`    // 用户 ID
	IsSeller  bool      `json:"is_seller"`  // 是否是对应卖家
	Content   string    `json:"content"`    // 评论内容
	CreatedAt time.Time `json:"created_at"` // 创建时间
	Username  string    `json:"username"`   // 用户名
	Nickname  string    `json:"nickname"`   // 昵称
	IconUrl   string    `json:"icon_url"`   // 用户头像
}

/////////////////////			基本增删改查			/////////////////////

// 根据回复的 ID 来获取回复条目
func FindCommentsReplyById(commentsReplyId uint64) (*CommentsReply, error) {
	var ret CommentsReply
	err := db.First(&ret, commentsReplyId).Error
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// 根据用户 ID 来获取回复列表
func FindCommentsReplysByUserId(userId uint64) ([]*CommentsReply, error) {
	var rets []*CommentsReply
	err := db.Where("user_id = ?", userId).Find(&rets).Error
	return rets, err
}

// 根据主评论 ID 来获取回复列表
func FindCommentsReplysByCommentId(commentId uint64) ([]*CommentsReply, error) {
	var rets []*CommentsReply
	err := db.Where("comment_id = ?", commentId).Find(&rets).Error
	return rets, err
}

// 新增回复
func InsertCommentsReply(commentsReply *CommentsReply) error {
	err := db.Create(commentsReply).Error
	return err
}

//删除回复条目 由于这里是软删除，所以只是把 delete_at 设置了一个值，实际上还存在数据库中,但并不能用 gorm 查到
func DeleteCommentsReplyById(id uint64) error {
	err := db.Delete(CommentsReply{}, "id = ?", id).Error
	return err
}

//根据 commentsReplyID 来更新其他相应的字段
func UpdateCommentsReply(commentsReply *CommentsReply) error {
	var rec CommentsReply
	err := db.First(&rec, "id = ?", commentsReply.ID).Error
	if err != nil {
		return err
	}
	err = db.Model(&rec).Updates(commentsReply).Error
	return err
}

/////////////////////		EMD	基本增删改查			/////////////////////
// 根据产品 ID 获取对应产品的评论
const sqlGetResultCommentReplaysByCommentId = `
 SELECT
 	c.*,u.username,u.nickname,u.icon_url
 FROM 
 comments_replys c JOIN
 users u ON c.user_id = u.id
 WHERE
 c.deleted_at IS NULL 
 AND u.deleted_at IS NULL 
 AND c.comment_id  = (?)
 ORDER BY
 id DESC
`

func GetResultCommentReplaysByCommentId(commentId uint64) ([]ResultCommentReply, error) {
	results := []ResultCommentReply{}
	err := db.Raw(sqlGetResultCommentReplaysByCommentId, commentId).Scan(&results).Error
	return results, err
}
