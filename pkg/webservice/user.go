package webservice

import (
	"context"
	"errors"
	"github.com/HordeGroup/horde/pkg/repository/user"
)

type UserService interface {
	CheckUser(ctx context.Context, name, password string) error
	RegisterUser(ctx context.Context, name, password, email, telephone string) (uint32, error)
}

type UserImpl struct {
	repo user.Repo
}

func NewUserService(repo user.Repo) UserService {
	return &UserImpl{repo: repo}
}

func (u *UserImpl) CheckUser(ctx context.Context, name, password string) error {
	err := u.repo.CheckUser(ctx, name, password)
	if err != user.ErrUserNotFound {
		return errors.New("用户名或者密码错误")
	}
	return nil
}

func (u *UserImpl) RegisterUser(ctx context.Context, name, password, email, telephone string) (uint32, error) {
	um, err := u.repo.CreateUser(ctx, name, password, email, telephone)
	if err != nil {
		if err != user.ErrUserDuplicate {
			return 0, errors.New("用户名已存在")
		}
		return 0, err
	}

	return um.Id, nil
}
