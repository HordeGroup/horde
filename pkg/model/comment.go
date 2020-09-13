package model

type Comment struct {
	Model
	// 评论内容
	Content string `gorm:"column:content;type:text;not null"`
	// 评论ID 若是跟帖回复,会存在跟帖的评论ID
	CommentId uint32 `gorm:"column:comment_id;type:int unsigned;default 0'"`
	// 帖子ID
	PostId uint32 `gorm:"column:post_id;type:int unsigned;not null"`
	// 用户ID 创建者ID
	CreatorId uint32 `gorm:"column:creator_id;type:int unsigned;not null"`
}

func (c Comment) TableName() string {
	return "comment"
}

func (c Comment) Indexes() map[string][]string {
	return map[string][]string{
		"creator_id": {"creator_id"},
		"comment_id": {"comment_id"},
	}
}
