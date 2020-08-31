package model

import "time"

type User struct {
	Id        uint32    `gorm:"id"`
	Name      string    `gorm: "name"`
	Password  string    `gorm: "password"`
	Salt      string    `gorm:"salt"`
	Email     string    `gorm:"email"`
	Telephone string    `gorm:"email"`
	CreateAt  time.Time `gorm:"creat_at"`
	UpdateAt  time.Time `gorm:"update_at"`
	Deleted   uint      `gorm:"deleted"`
}
