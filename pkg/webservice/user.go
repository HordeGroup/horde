package webservice

import (
	"context"
	"github.com/jinzhu/gorm"
)

type UserService interface {
	CheckUser(ctx context.Context, name, password string) error
	CreateUser(ctx context.Context, name, password, telephone string) (uint32, error)
}

type UserImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &UserImpl{db: db}
}

func (u *UserImpl) CheckUser(ctx context.Context, name, password string) error {
	panic("implement me")
}

func (u *UserImpl) CreateUser(ctx context.Context, name, password, telephone string) (uint32, error) {
	panic("implement me")
}
