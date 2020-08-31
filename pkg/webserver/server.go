package webserver

import (
	"github.com/rs/zerolog"
)

type Option struct {
	Host   string
	Port   int
	Logger zerolog.Logger
}

type Server struct {
	host   string
	port   int
	logger zerolog.Logger
}

func New() *Server {
	return &Server{}
}

func Run() error {
	return nil
}
