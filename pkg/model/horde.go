package model

/**
贴吧
*/
type Horde struct {
	Model
	Id uint32 `gorm:"column:id;primary_key;auto_increment"`
	// 名称
	Name string `gorm:"column:name;type:char(16);not null"`
	// 关注
	Focus int64 `gorm:"column:focus;type:int;not null;default:0"`
	//// 描述主体
	//Content string `gorm:"column:content;type:varchar(255)"`
	// 简介
	Description string `gorm:"column:discription;type:varchar(100)"`
	// 创建者
	CreatorId uint32 `gorm:"column:creator_id;type:int unsigned;not null"`
}

func (h Horde) TableName() string {
	return "horde"
}

func (h Horde) UniqueIndexes() map[string][]string {
	return map[string][]string{
		"name_dt": {"name", "deleted"},
	}
}

func (h Horde) Indexes() map[string][]string {
	return map[string][]string{
		"creator": {"creator_id"},
	}
}
