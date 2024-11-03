// main.go
package main

import (
	"go-crud/config"
	"go-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDB()

	router := gin.Default()

	// add a ping route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.RegisterRoutes(router)

	// Run server
	router.Run(":3000")
}
