package user

import (
	"context"
	"github.com/HordeGroup/horde/pkg/model"
	"github.com/jinzhu/gorm"
)

type Repo interface {
	CheckUser(ctx context.Context, name, password string) error
	CreateUser(ctx context.Context, name, password, email, telephone string) (model.User, error)
}

func NewRepo(db *gorm.DB) Repo {
	return &repoImpl{db: db}
}
