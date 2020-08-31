package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Option struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Debug    bool
}

func New(opt Option) (*gorm.DB, error) {
	connectUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&%s",
		opt.User, opt.Password, opt.Host, opt.Port, opt.Name, "clientFoundRows=true")
	db, err := gorm.Open("mysql", connectUrl)
	if err != nil {
		return nil, err
	}
	db.LogMode(opt.Debug)
	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(time.Hour)
	return db, nil
}
