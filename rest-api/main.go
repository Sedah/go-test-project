package main

import (
	"github.com/Sedah/go-test-project/controllers"
	"github.com/Sedah/go-test-project/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	r.GET("/books", controllers.FindBooks)
	r.POST("/book", controllers.CreateBook)
	r.GET("/book/:id", controllers.FindBook)
	r.PATCH("/book/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run()
}
