package model

import "time"

type User struct {
	Id        uint32    `gorm:"column:id;primary_key;auto_increment"`
	Name      string    `gorm:"column:name;type:char(16);not null"`
	Password  string    `gorm:"column:password;type:char(16);not null"`
	Salt      string    `gorm:"column:salt;type:char(32);not null"`
	Email     string    `gorm:"column:email;type:varchar(255)"`
	Telephone string    `gorm:"column:telephone;type:varchar(255)"`
	CreateAt  time.Time `gorm:"column:created_at;type:timestamp(6);not null;default:now(6)"`
	UpdateAt  time.Time `gorm:"column:update_at;type:timestamp(6);not null;default:now(6)"`
	Deleted   uint      `gorm:"column:deleted;type:tinyint;not null;default:0"`
}

func (User) TableName() string {
	return "user"
}

func (User) UniqueIndexes() map[string][]string {
	return map[string][]string{
		"name": {"name", "deleted"},
	}
}

func (User) Indexes() map[string][]string {
	return map[string][]string{
		"email":     {"email"},
		"telephone": {"telephone"},
	}
}
