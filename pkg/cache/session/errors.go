package session

import "github.com/pkg/errors"

var (
	ErrSessionTokenNotFound   = errors.New("session token not found")
	ErrDuplicatedSessionToken = errors.New("duplicate session token")
)
