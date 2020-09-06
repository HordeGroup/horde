package webserver

import (
	"github.com/HordeGroup/horde/pkg/cache/session"
	"github.com/HordeGroup/horde/pkg/def"
	"github.com/HordeGroup/horde/pkg/render"
	"github.com/gin-gonic/gin"
)

func (sv *Server) UserLogin(c *gin.Context) {
	var (
		userLogin def.UserLoginRequest
		sess      session.Session
		err       error
	)
	if err = c.ShouldBind(&userLogin); err != nil {
		render.JSONWithError(c, err)
		return
	}
	sess, err = sv.service.User.Login(c.Request.Context(), userLogin.Name, userLogin.Password)
	if err != nil {
		render.JSONWithError(c, err)
		return
	}
	c.SetCookie("session_id", sess.Token, 1, "", "localhost", true, true)
	render.JSONSuccess(c, def.UserLoginData{UserId: sess.UserId})
}

func (sv *Server) UserRegister(c *gin.Context) {
	var userRegister def.UserRegisterRequest
	if err := c.ShouldBind(&userRegister); err != nil {
		render.JSONWithError(c, err)
		return
	}
	userId, err := sv.service.User.Register(c.Request.Context(), userRegister.Name, userRegister.Password, userRegister.Email, userRegister.Telephone)
	if err != nil {
		render.JSONWithError(c, err)
		return
	}
	render.JSONSuccess(c, def.UserRegisterData{UserId: userId})
}
