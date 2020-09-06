package webserver

import (
	"github.com/HordeGroup/horde/pkg/def"
	"github.com/HordeGroup/horde/pkg/render"
	"github.com/gin-gonic/gin"
)

func (sv *Server) UserLogin(c *gin.Context) {
	var (
		userLogin def.UserLoginRequest
		err       error
	)
	if err = c.ShouldBind(&userLogin); err != nil {
		render.JSONWithError(c, err)
		return
	}
	err = sv.service.User.CheckUser(c.Request.Context(), userLogin.Name, userLogin.Password)
	if err != nil {
		render.JSONWithError(c, err)
		return
	}
	render.JSONSuccess(c, def.UserLoginData{})
}

func (sv *Server) UserRegister(c *gin.Context) {
	var userRegister def.UserRegisterRequest
	if err := c.ShouldBind(&userRegister); err != nil {
		render.JSONWithError(c, err)
		return
	}
	userId, err := sv.service.User.RegisterUser(c.Request.Context(), userRegister.Name, userRegister.Password, userRegister.Email, userRegister.Telephone)
	if err != nil {
		render.JSONWithError(c, err)
		return
	}
	render.JSONSuccess(c, def.UserRegisterData{UserId: userId})
}
