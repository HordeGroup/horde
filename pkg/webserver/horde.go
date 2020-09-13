package webserver

import (
	"github.com/HordeGroup/horde/pkg/def"
	"github.com/HordeGroup/horde/pkg/herror"
	"github.com/HordeGroup/horde/pkg/model"
	"github.com/HordeGroup/horde/pkg/render"
	"github.com/gin-gonic/gin"
)

func (sv *Server) CreateHorde(c *gin.Context) {
	var (
		ctx            = c.Request.Context()
		createHordeReq def.HordeCreateRequest
		hm             model.Horde
		err            error
	)
	if err = c.ShouldBind(&createHordeReq); err != nil {
		render.JSONWithError(c, herror.ErrInvalidRequest)
		return
	}
	if err = sv.service.User.Exists(ctx, createHordeReq.CreatorId); err != nil {
		render.JSONWithError(c, herror.ErrUserNotFound)
		return
	}
	if hm, err = sv.service.Horde.CreateHorde(c.Request.Context(), createHordeReq.Name, createHordeReq.Desc, createHordeReq.CreatorId); err != nil {
		render.JSONWithError(c, herror.ErrCreateHorde)
		return
	}
	render.JSONSuccess(c, &def.HordeCreateData{Id: hm.Id})
}
