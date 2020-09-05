package user

import (
	"context"
	"github.com/HordeGroup/horde/pkg/model"
	"github.com/jinzhu/gorm"
)

type repoImpl struct {
	db *gorm.DB
}

func (r *repoImpl) CheckUser(ctx context.Context, name, password string) error {
	um := model.User{
		Name:     name,
		Password: password,
	}
	return r.db.Find(um).Error
}

func (r *repoImpl) CreateUser(ctx context.Context, name, password, email, telephone string) (model.User, error) {
	um := model.User{
		Name:      name,
		Password:  password,
		Salt:      telephone,
		Email:     email,
		Telephone: telephone,
	}
	err := r.db.Create(um).Error
	if err != nil {
		return um, nil
	}
	return um, nil
}
