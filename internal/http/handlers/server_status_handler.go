package handler

import "github.com/gin-gonic/gin"

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"The server is up and running",
		"status": "ok",
	})
}
