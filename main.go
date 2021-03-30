package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"onego/apis"
	"onego/int"
)

// Index returns a home page
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"projetc_name": "test", "version": "v1", "url": "/v1"})
}

func main() {
	int.Mysql()
	int.Redis()
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", Index)
	v1 := router.Group("v1")
	v1.GET("/novels/:name/:source", apis.SearchNovels)
	v1.GET("/authors/:name/:source", apis.SearchAuthors)
	v1.GET("/book/:bookid/", apis.bookview)
	v1.GET("/read/:bookid/:chapterid", apis.chapterview)
	router.Run()
}
