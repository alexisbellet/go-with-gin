package handlers

import (
	"go-with-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
var albums = []models.Album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
    // albums slice to seed record album data.
    c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusFound, a) 
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbums(c *gin.Context) {
    var newAlbum models.Album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
    
    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}