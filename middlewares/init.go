package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pierorolando1/crudmongo/helpers"
)

func ValidateToken(c *gin.Context) {

	//get the bearer token
	token := c.GetHeader("Authorization")

	//check if the token is valid
	valid, err := helpers.ValidateToken(token)

	if err != nil || !valid {
		c.JSON(400, gin.H{
			"ok":    false,
			"error": "invalid token in middleware",
		})
		c.Abort()
		return
	}

	c.Next()
}
