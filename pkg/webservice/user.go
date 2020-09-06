package webservice

import (
	"context"
	"github.com/HordeGroup/horde/pkg/cache/session"
	"github.com/HordeGroup/horde/pkg/helper"
	"github.com/HordeGroup/horde/pkg/herror"
	"github.com/HordeGroup/horde/pkg/model"
	"github.com/HordeGroup/horde/pkg/repository/user"
	zlog "github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, name, password string) (session.Session, error)
	Register(ctx context.Context, name, password, email, telephone string) (uint32, error)
}

type UserImpl struct {
	repo      user.Repo
	sessCache session.Cache
}

func NewUserService(repo user.Repo, sessCache session.Cache) UserService {
	return &UserImpl{repo: repo, sessCache: sessCache}
}

func (u *UserImpl) Login(ctx context.Context, name, password string) (sess session.Session, err error) {
	var (
		um model.User
	)
	if um, err = u.repo.GetUserByName(ctx, name); err != nil {
		err = herror.ErrInvalidNameOrPwd
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(um.Password), []byte(password)); err != nil {
		err = herror.ErrInvalidNameOrPwd
		return
	}
	if sess, err = u.sessCache.New(um.Id); err != nil {
		zlog.Err(err).Msg("创建SESSION失败")
		return
	}

	return
}

func (u *UserImpl) Register(ctx context.Context, name, password, email, telephone string) (uint32, error) {
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
