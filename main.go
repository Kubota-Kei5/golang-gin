package main

import (
	"github.com/gin-gonic/gin"

	"golang-gin/controllers"
	"golang-gin/models"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", controllers.Ping)

	albumHandler := controllers.AlbumHandler{}
	router.POST("/album", albumHandler.CreateAlbum)
	router.GET("/album/:id", albumHandler.GetAlbum)
	router.DELETE("/album/:id", albumHandler.DeleteAlbum)
	return router
}

func main() {
	db, _ := models.ConnectToDB()
	models.DB = db

	albumModels := []interface{}{&models.Album{}}
	for _, val := range albumModels {
		models.DB.AutoMigrate(val)
	}

	router := setupRouter()
	router.Run(":8080")
}