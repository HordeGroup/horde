package webserver

import (
	"github.com/HordeGroup/horde/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
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
		url := ginswagger.URL("http://localhost:9000/swagger/doc.json")
		engine.GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler, url))
	}

	{
		engine.POST("/user", sv.UserRegister)
		engine.POST("/user/login", sv.UserLogin)
	}

	{
		engine.POST("/horde", sv.CreateHorde)
	}

	return engine
}
