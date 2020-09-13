package herror

import "github.com/pkg/errors"

var (
	ErrInvalidRequest = errors.New("请求参数错误")

	ErrInvalidNameOrPwd     = errors.New("用户名或者密码错误")
	ErrInvalidUserName      = errors.New("用户名不合法")
	ErrInvalidUserPwd       = errors.New("用户密码不合法")
	ErrInvalidUserEmail     = errors.New("用户邮箱不合法")
	ErrInvalidUserTelephone = errors.New("用户电话不合法")
	ErrUserAlreadyExists    = errors.New("用户已存在")
	ErrUserNotFound         = errors.New("用户不存在")

	ErrCreateHorde = errors.New("创建部落失败")
)
