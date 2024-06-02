package main

import (
	"github.com/egespindola/goapi-ws-gin/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", service.GetAlbums)
	router.GET("/albums/:id", service.GetAlbumByID)
	router.POST("/albums", service.AddAlbums)

	router.Run("localhost:8080")
}
