package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/egespindola/goapi-ws-gin/internal/database"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Albums)
}

// post h service
func AddAlbums(c *gin.Context) {
	var newAlbum database.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	database.Albums = append(database.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// get specific item
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range database.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
