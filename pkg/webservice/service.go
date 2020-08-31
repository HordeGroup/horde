package webservice

import "github.com/jinzhu/gorm"

type Option struct {
	Verbose bool
	DB      *gorm.DB
}

type Service struct {
	verbose bool
	User    UserService
}

func New(opt Option) *Service {
	return &Service{
		verbose: opt.Verbose,
		User:    NewUserService(opt.DB),
	}
}
