// Package server implements the HTTP server using the Gin framework.
package server

import (
	"fmt"

	"gopastbin/internal/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	port   string
}

// NewServer initializes the Gin router and returns a Server instance
func NewServer(port string) *Server {
	r := http.SetupRoutes()

	// setup all routes here
	return &Server{
		router: r,
		port:   port,
	}
}

// Start runs the server
func (s *Server) Start() error {
	fmt.Println("Server running on port", s.port)
	return s.router.Run(":" + s.port)
}
