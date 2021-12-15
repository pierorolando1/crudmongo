package routes

import (
	"github.com/gin-gonic/gin"
)

var R = gin.Default()

func Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the CRUD API",
	})
}
