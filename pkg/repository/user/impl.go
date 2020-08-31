package user

import (
	"context"
	"github.com/jinzhu/gorm"
)

type repoImpl struct {
	db *gorm.DB
}

func (r *repoImpl) CheckUser(ctx context.Context, name, password string) error {
	panic("implement me")
}

func (r *repoImpl) CreateUser(ctx context.Context, name, password, telephone string) (uint32, error) {
	panic("implement me")
}
