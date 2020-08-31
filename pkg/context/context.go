package context

import (
	"github.com/gin-gonic/gin/render"
	"github.com/rs/zerolog"
)

type Context interface {
	Id() string
	Logger() zerolog.Logger
	Type() string
	Auth() interface{}
	Render() render.Render
}
