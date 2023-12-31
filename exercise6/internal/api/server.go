package api

import (
	"github.com/gin-gonic/gin"
	"github.com/diegovanne/go23/exercise6/internal/api/routes"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{
		router: gin.Default(),
	}

	routes.InitializeRoutes(server.router)

	return server
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
