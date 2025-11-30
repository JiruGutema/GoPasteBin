// Package handler handles HTTP requests and responses
package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Println("Login handler called")
	c.JSON(200, gin.H{
		"message": "Login Success!",
		"token":   "thisisfaketoken",
		"user":    "Jiru Gutema",
	})
}

func Register(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "Registered successfully!",
		"token":   "thisisfaketoken",
		"user":    "Jiru Gutema",
	})
}
