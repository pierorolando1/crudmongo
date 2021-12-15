package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pierorolando1/crudmongo/routes"
)

func main() {

	routes.R.GET("/", routes.Welcome)
	routes.R.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//this executes the auth routes
	routes.AuthRoutes()

	err := routes.R.Run()
	if err != nil {
		panic(err)
	}

	println("Server is running on port 8080")
}
