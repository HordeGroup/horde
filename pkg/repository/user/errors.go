package user

import "github.com/pkg/errors"

var (
	ErrUserNotFound  = errors.New("User Err: Not Found")
	ErrUserDuplicate = errors.New("User Err: already exists")
)
