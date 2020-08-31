package webserver

import (
	"github.com/gin-gonic/gin"
)

type UserRegisterRequest struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}

type UserLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (sv *Server) UserLogin(c *gin.Context) {

}

func (sv *Server) UserRegister(c *gin.Context) {

}
