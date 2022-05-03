package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jpaola/web-service-gin/records"
)

func main() {

	// initialize gin router using Default
    router := gin.Default()

	router.GET("/albums", records.GetAlbums)
    router.GET("/albums/:id", records.GetAlbumByID)
    router.POST("/albums", records.AddAlbum)
    router.DELETE("/albums/:id", records.RemoveAlbumByID)
    router.PUT("/albums/:id", records.UpdateAlbumByID)

    router.Run("localhost:8080")
}