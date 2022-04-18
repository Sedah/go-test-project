package main

import (
	"github.com/Sedah/go-test-project/rest-api-2/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.POST("/albums", controllers.PostAlbums)
	router.PATCH("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)

	router.Run("localhost:8080")
}
