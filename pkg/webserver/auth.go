package webserver

import (
	"github.com/HordeGroup/horde/pkg/def"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (sv *Server) AuthWrapper() gin.HandlerFunc {

	return func(c *gin.Context) {
		var (
			sessionId string
			err       error
		)
		if sessionId, err = c.Cookie("session_id"); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, def.NeedLoginData{NeedLogin: true})
			return
		}
		if _, err = sv.service.SessCache.Get(sessionId); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, def.NeedLoginData{NeedLogin: true})
			return
		}
		c.Next()
	}

}
