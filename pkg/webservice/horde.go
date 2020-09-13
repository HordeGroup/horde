package webservice

import (
	"context"
	"github.com/HordeGroup/horde/pkg/model"
	"github.com/HordeGroup/horde/pkg/repository/horde"
)

type HordeService interface {
	CreateHorde(ctx context.Context, name string, desc string, creatorId uint32) (model.Horde, error)
}

type hordeImpl struct {
	repo horde.Repo
}

func (h *hordeImpl) CreateHorde(ctx context.Context, name string, desc string, creatorId uint32) (hm model.Horde, err error) {
	if hm, err = h.repo.Create(ctx, name, desc, creatorId); err != nil {
		return
	}
	return
}

func NewHordeService(repo horde.Repo) HordeService {
	return &hordeImpl{repo: repo}
}
