package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	default_port = "8080"
	HTTP_PORT    = "PORT"
)

type Server struct {
	port   string
	router *gin.Engine
}

func New(p string) *Server {
	httpPort := ":"
	if p == "" {
		if env, ok := os.LookupEnv(HTTP_PORT); !ok {
			httpPort += default_port
		} else {
			httpPort += env
		}
	} else {
		httpPort += p
	}
	return &Server{
		port:   httpPort,
		router: gin.New(),
	}
}

func (s *Server) Run() {
	log.Panic(s.router.Run(":8080"))
}
