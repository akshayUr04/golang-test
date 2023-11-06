package api

import (
	"golang-test/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewHttpServer() *Server {
	engine := gin.New()
	engine.POST("/", handler.GetRequest)
	return &Server{
		engine: engine,
	}
}

func (s *Server) Start() {
	s.engine.Run(":8001")
}
