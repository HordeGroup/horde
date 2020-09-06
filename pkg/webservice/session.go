package webservice

import "time"

type Session struct {
	Value      string
	ExpireTime time.Time
}

type SessionService interface {
	Get(session Session) error
	Save(session Session) error
	Delete(session Session) error
}

type sessionServiceImpl struct {
}
