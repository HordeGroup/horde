package webserver

import (
	"github.com/HordeGroup/horde/pkg/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (sv *Server) BuildRouter() *gin.Engine {
	engine := gin.New()

	if sv.verbose {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine.Use(middleware.RequestDump(sv.verbose, sv.logger))

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	{
		engine.POST("/user", sv.UserRegister)
		engine.POST("/user/login", sv.UserLogin)
	}

	return engine
}
