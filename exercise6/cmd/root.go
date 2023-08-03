package cmd

import (
	"log"
	"github.com/diegovanne/go23/exercise6/internal/api"
)

func Execute() {
	server := api.NewServer()
	if err := server.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}