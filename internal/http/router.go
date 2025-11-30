// Package http contains all of the router endpoints mapped to the their hanlder functions
package http

import (
	handler "gopastbin/internal/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/health", handler.Health)
	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)

	return r
}
