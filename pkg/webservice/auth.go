package webservice

import "github.com/gin-gonic/gin"

type Auth interface {
	GetAuth() (interface{}, error)

	OnError(c *gin.Context, err error)

	WriteAuth() (c *gin.Context)
}
