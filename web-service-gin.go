package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  string `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: "$56.99"},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: "$17.99"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: "$39.99"},
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "album not found",
	})
}

func postAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "album deleted",
			})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "album not found",
	})
}

func updateAlbum(c *gin.Context) {
	var albumUpdate Album

	if err := c.BindJSON(&albumUpdate); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}
	for i, album := range albums {
		if album.ID == albumUpdate.ID {
			albums[i] = albumUpdate
			c.IndentedJSON(http.StatusOK, albumUpdate)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "album not found",
	})
}

func webServiceGin() {
	r := gin.Default()
	r.GET("/health", healthCheck)
	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbumByID)
	r.POST("/albums", postAlbums)
	r.DELETE("/albums/:id", deleteAlbum)
	r.PUT("/albums", updateAlbum)

	r.Run(":8080")
}
