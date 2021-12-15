package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pierorolando1/crudmongo/controllers"
	"github.com/pierorolando1/crudmongo/middlewares"
)

func AuthRoutes() {
	R.POST("/login", controllers.Logincontroller)

	R.POST("/register", controllers.Registercontroller)

	R.GET("/hola", middlewares.ValidateToken, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ok": true,
		})
	})
}
