package webservice

import (
	"context"
	"github.com/HordeGroup/horde/pkg/helper"
	"github.com/HordeGroup/horde/pkg/herror"
	"github.com/HordeGroup/horde/pkg/model"
	"github.com/HordeGroup/horde/pkg/repository/user"
	zlog "github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
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
	var (
		um  model.User
		err error
	)
	if um, err = u.repo.GetUserByName(ctx, name); err != nil {
		return herror.ErrInvalidNameOrPwd
	}
	if err = bcrypt.CompareHashAndPassword([]byte(um.Password), []byte(password)); err != nil {
		return herror.ErrInvalidNameOrPwd
	}
	return nil
}

func (u *UserImpl) RegisterUser(ctx context.Context, name, password, email, telephone string) (uint32, error) {
	if !helper.CheckUserName(name) {
		return 0, herror.ErrInvalidUserName
	}
	if !helper.CheckUserPwd(password) {
		return 0, herror.ErrInvalidUserPwd
	}

	pwdHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, herror.ErrInvalidUserPwd
	}
	um, err := u.repo.CreateUser(ctx, name, string(pwdHashed), email, telephone)
	if err != nil {
		zlog.Err(err).Msg("创建用户失败")
		return 0, herror.ErrUserAlreadyExists
	}

	return um.Id, nil
}
