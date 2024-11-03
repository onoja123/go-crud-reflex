package controllers

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks retrieves all books from the database.
func GetBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)

	if len(books) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No books found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GetOneBook retrieves a single book by its ID.
func GetOneBook(c *gin.Context) {
	var book models.Book

	var id = c.Param("id")

	config.DB.First(&book, id)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook adds a new book to the database.
func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&book)
	c.JSON(http.StatusCreated, gin.H{"data": book})
}

// DeleteBook removes a book from the database by its ID.
func DeleteBook(c *gin.Context) {
	var book models.Book

	if err := config.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	config.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
