package webservice

import (
	"github.com/HordeGroup/horde/pkg/cache/session"
	"github.com/HordeGroup/horde/pkg/repository/user"
	"github.com/jinzhu/gorm"
)

type Option struct {
	Verbose   bool
	DB        *gorm.DB
	SessCache session.Cache
}

type Service struct {
	verbose   bool
	User      UserService
	SessCache session.Cache
}

func New(opt Option) *Service {
	return &Service{
		verbose: opt.Verbose,
		User:    NewUserService(user.NewRepo(opt.DB), opt.SessCache),
	}
}
