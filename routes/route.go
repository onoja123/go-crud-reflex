package routes

import (
	"github.com/gin-gonic/gin"
	"go-crud/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetOneBook)
	router.POST("/books", controllers.CreateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
}
