package database

import (
	"fmt"
	"github.com/HordeGroup/horde/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/gormigrate.v1"
	"time"
)

type Option struct {
	Name     string
	Host     string
	Port     int64
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

func GetMigration() *Migration {
	return &Migration{
		Options: gormigrate.DefaultOptions,
		InitSchema: func(db *gorm.DB) error {
			return MigrateTables(db, &model.User{}).Error
		},
		Migrations: nil,
	}
}
