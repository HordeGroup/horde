package user

import (
	"context"
	"github.com/HordeGroup/horde/pkg/model"
	"github.com/jinzhu/gorm"
)

type repoImpl struct {
	db *gorm.DB
}

func (r *repoImpl) GetUserByName(ctx context.Context, name string) (um model.User, err error) {
	err = r.db.Where("name = ?", name).Find(&um).Error
	if gorm.IsRecordNotFoundError(err) {
		err = ErrUserNotFound
	}
	return
}

func (r *repoImpl) CheckUser(ctx context.Context, name, password string) (err error) {
	um := model.User{
		Name:     name,
		Password: password,
	}
	err = r.db.Where("name = ?", um.Name).Where("password = ?", password).Find(&um).Error
	if gorm.IsRecordNotFoundError(err) {
		err = ErrUserNotFound
	}
	return
}

func (r *repoImpl) CreateUser(ctx context.Context, name, password, email, telephone string) (model.User, error) {
	um := model.User{
		Name:      name,
		Password:  password,
		Email:     email,
		Telephone: telephone,
	}
	err := r.db.Create(&um).Error
	if err != nil {
		return um, err
	}
	return um, nil
}

func (r *repoImpl) GetById(ctx context.Context, id uint32) (um model.User, err error) {
	if err = r.db.Table(um.TableName()).
		Where("id = ? ", id).
		First(&um).Error; err != nil {
		return
	}
	return
}
