package horde

import (
	"context"
	"github.com/HordeGroup/horde/pkg/model"
	"github.com/jinzhu/gorm"
)

type Repo interface {
	Create(ctx context.Context, name string, desc string, creatorId uint32) (model.Horde, error)
	UpdateDesc(ctx context.Context, desc string) error
	Delete(ctx context.Context, id uint32) error
}

func NewRepo(db *gorm.DB) Repo {
	return &impl{db: db}
}
