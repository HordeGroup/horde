package webservice

import (
	"github.com/HordeGroup/horde/pkg/cache/session"
	"github.com/HordeGroup/horde/pkg/repository/horde"
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
	SessCache session.Cache
	User      UserService
	Horde     HordeService
}

func New(opt Option) *Service {
	return &Service{
		verbose:   opt.Verbose,
		SessCache: opt.SessCache,
		User:      NewUserService(user.NewRepo(opt.DB), opt.SessCache),
		Horde:     NewHordeService(horde.NewRepo(opt.DB)),
	}
}
