package controllers

import (
	"net/http"

	"github.com/Sedah/go-test-project/rest-api-2/models"
	"github.com/gin-gonic/gin"
)

type UpdateAlbumInput struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, models.Albums)
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new album to the slice.
	models.Albums = append(models.Albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range models.Albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// Update a book
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")

	// Validate input
	var input UpdateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update album
	for i, a := range models.Albums {
		if a.ID == id {
			models.Albums[i].Title = input.Title
			models.Albums[i].Artist = input.Artist
			models.Albums[i].Price = input.Price
		}
	}

	c.JSON(http.StatusOK, models.Albums)
}

// Delete a book
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	for i, a := range models.Albums {
		if a.ID == id {
			models.Albums = append(models.Albums[0:i], models.Albums[i+1:]...)
		}
	}

	c.JSON(http.StatusOK, models.Albums)
}
