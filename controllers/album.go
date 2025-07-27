package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"golang-gin/models"
)

type AlbumHandler struct{}

func (a *AlbumHandler) CreateAlbum(c *gin.Context) {
	var album models.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	createdAlbum, err := album.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, *createdAlbum)
}

func (a *AlbumHandler) GetAlbum(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	getAlbum, err := models.AlbumFindOne(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, getAlbum)
}

func (a *AlbumHandler) DeleteAlbum(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	album, _ := models.AlbumFindOne(ID)
	if err := album.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}