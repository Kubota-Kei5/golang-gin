package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID	 string  `json:"id"`
	Title string  `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Album 1", Artist: "Artist 1", Price: 9.99},
	{ID: "2", Title: "Album 2", Artist: "Artist 2", Price: 14.99},
	{ID: "3", Title: "Album 3", Artist: "Artist 3", Price: 19.99},
}


func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(context *gin.Context) {
	var newAlbum album

	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(context *gin.Context) {
	id := context.Param("id")
	for _, album := range albums {
		if album.ID == id {
			context.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


func main() {
	r :=  gin.Default()
	r.GET("/albums", getAlbums)
	r.POST("/albums", postAlbums)
	r.GET("/albums/:id", getAlbumByID)
	r.Run("localhost:8080")
}