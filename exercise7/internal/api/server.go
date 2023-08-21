package api

import (
	"github.com/gin-gonic/gin"
	"github.com/diegovanne/go23/exercise7/internal/api/routes"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{
		router: gin.Default(),
	}

	routes.InitializeRoutes(server.router)s

	return server
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
