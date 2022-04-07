package main

import (
	"go-with-gin/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := gin.Default()
    router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
    router.POST("/albums", handlers.PostAlbums)

    router.Run("localhost:8080")
}