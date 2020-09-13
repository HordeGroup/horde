package horde

import (
	"github.com/HordeGroup/horde/pkg/context"
	"github.com/HordeGroup/horde/pkg/model"
	"github.com/jinzhu/gorm"
)

type impl struct {
	db *gorm.DB
}

func (i *impl) Create(ctx context.Context, name string, desc string, creatorId uint32) (hm model.Horde, err error) {
	hm = model.Horde{
		Id:          0,
		Name:        name,
		Description: desc,
		CreatorId:   creatorId,
	}
	if err = i.db.Create(&hm).Error; err != nil {
		return
	}
	return
}

func (i *impl) UpdateDesc(ctx context.Context, desc string) (err error) {
	if err = i.db.Table("horde").UpdateColumn("desc", desc).Error; err != nil {
		return
	}
	return
}

func (i *impl) Delete(ctx context.Context, id uint32) (err error) {
	err = i.db.Delete(&model.Horde{}, id).Error
	return
}
