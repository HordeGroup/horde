package webserver

import "github.com/gin-gonic/gin"

func (sv *Server) AuthWrapper() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()
	}

}
