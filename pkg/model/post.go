package model

/**
帖子
*/
type Post struct {
	Model
	// 主题
	Title string `gorm:"column:title;type:varchar(100);not null"`
	// 描述
	Description string `gorm:"column:description;type:varchar(256)"`
	// 创建者
	CreatorId uint32 `gorm:"column:creator_id;type:int unsigned;not null"`
}

func (p Post) TableName() string {
	return "post"
}

func (p Post) Indexes() map[string][]string {
	return map[string][]string{
		"creator_id": {"creator_id"},
	}
}
