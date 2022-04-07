package handlers

import (
	"database/sql"
	"fmt"
	"go-with-gin/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func connectToDb() *sql.DB {
    // Open a connection to the database
	connStr := fmt.Sprintf("host=%v dbname=%v user=%v password=%v sslmode=%v", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_SSLMODE"))
    log.Print(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not connect to the database: %v", err))
	}
    return db
}

func GetAlbums(c *gin.Context) {
    db := connectToDb()

    // An albums slice to hold data from returned rows.
    var albums []models.Album

    // Query the db for albums
    rows, err := db.Query(`SELECT * FROM albums`)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, err)
    }

    defer rows.Close()
    
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb models.Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            c.IndentedJSON(http.StatusInternalServerError, err)
            return
        }
        albums = append(albums, alb)
    }

    if err := rows.Err(); err != nil {
        c.IndentedJSON(http.StatusInternalServerError, err)
        return
    }

    c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
    db := connectToDb()

    // An album struct to hold data from returned rows.
    var alb models.Album
    id := c.Param("id")

    // Query the db for albums
    row := db.QueryRow(`SELECT * FROM albums WHERE id = $1`, id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
        if err == sql.ErrNoRows {
            c.IndentedJSON(http.StatusNotFound, fmt.Sprintf("Could not find album with id: %v", id))
            return
        }
        c.IndentedJSON(http.StatusInternalServerError, err)
        return
    }
    c.IndentedJSON(http.StatusFound, alb)
}

// func PostAlbums(c *gin.Context) {
//     var newAlbum models.Album

//     if err := c.BindJSON(&newAlbum); err != nil {
//         return
//     }
    
//     // Add the new album to the slice.
//     albums = append(albums, newAlbum)
//     c.IndentedJSON(http.StatusCreated, newAlbum)
// }