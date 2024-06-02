package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/egespindola/goapi-ws-gin/internal/database"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Albums)
}

func AddAlbums(c *gin.Context) {
	var newAlbum database.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	database.Albums = append(database.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
