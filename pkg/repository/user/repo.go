package user

import "context"

type Repo interface {
	CheckUser(ctx context.Context, name, password string) error
	CreateUser(ctx context.Context, name, password, telephone string) (uint32, error)
}
