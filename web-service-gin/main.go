package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// return with list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// add an album from JSON received in request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// bind received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// return target Album by given ID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// delete Album by given ID
func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	leftAlbums := []album{}
	var removedAlbum album

	for _, a := range albums {
		if a.ID == id {
			removedAlbum = a
			continue
		}
		leftAlbums = append(leftAlbums, a)
	}
	albums = leftAlbums

	if removedAlbum.ID != "" {
		c.IndentedJSON(http.StatusOK, removedAlbum)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	//router := gin.New()
	// Default with the logger and recovery middleware already attached
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	//grouping routes
	v2 := router.Group("/v2")
	{
		v2.POST("albums", postAlbums)
		v2.GET("/albums", getAlbums)
		v2.GET("/album/:id", func(c *gin.Context) {
			c.String(http.StatusMethodNotAllowed, "this API is not allowed")
		})
	}

	// // Authorization group
	// authorized := router.Group("/")
	// authorized.Use(AuthRequired())
	// {
	// 	authorized.POST("")
	// }

	router.Run("localhost:8080")
}
