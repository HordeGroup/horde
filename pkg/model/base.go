package model

import "time"

type Model struct {
	CreateAt time.Time `gorm:"column:created_at;type:timestamp(6);not null;default:now(6)"`
	UpdateAt time.Time `gorm:"column:update_at;type:timestamp(6);not null;default:now(6)"`
	Deleted  uint      `gorm:"column:deleted;type:tinyint;not null;default:0"`
}
