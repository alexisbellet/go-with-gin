package main

import (
	. "go-with-gin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
    router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumByID)
    router.POST("/albums", PostAlbums)

    router.Run("localhost:8080")
}