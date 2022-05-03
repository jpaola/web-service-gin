package records

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpaola/web-service-gin/database"
	entity "github.com/jpaola/web-service-gin/entity"
)

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(ctx *gin.Context) {
	db := database.DBConn()
	albums := []entity.Album{}

	results, err := db.Query("SELECT * FROM albums")

	// if there is an error during query, handle it
	if err != nil {
		fmt.Println("An error occured while attempting to fetch albums data", err.Error())
		return
	}

	for results.Next() {
		var a entity.Album
		// for each row, scan into the album struct
		err = results.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)

		if err != nil {
			fmt.Println("Error while scanning album struct", err.Error()) // proper error handling instead of panic in your app
		}
        // append the album into albums array
		albums = append(albums, a)
	}

    ctx.IndentedJSON(http.StatusOK, albums)
	defer db.Close()
}

// AddAlbum adds an album from JSON received in the request body.
func AddAlbum(ctx *gin.Context) {
	db := database.DBConn()
	var newAlbum entity.Album

	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insert, err := db.Query(
		"INSERT INTO albums (albumId,title,artist,price) VALUES (?,?,?,?)",
		newAlbum.ID, newAlbum.Title, newAlbum.Artist, newAlbum.Price)

	// if there is an error inserting, handle it
	if err != nil {
		fmt.Println("An error occured while attempting to insert new album into db", err.Error())
	}

	defer insert.Close()
 	ctx.JSON(http.StatusOK, gin.H{"data": insert})
	defer db.Close()
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(ctx *gin.Context) {
	db := database.DBConn()
	id := ctx.Param("id")
	album := &entity.Album{}

	newAlbum, err := db.Query("SELECT * FROM albums where albumId=?", id)

	// if there is an error during query, handle it
	if err != nil {
		fmt.Println("An error occured while attempting to fetch album from db", err.Error())
		return
	}
	 
	if newAlbum.Next() {
		err = newAlbum.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			fmt.Println("Error while scanning album struct", err.Error()) // proper error handling instead of panic in your app
		} 
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

	ctx.IndentedJSON(http.StatusOK, album)
	defer db.Close()
}

// RemoveAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then removes that album as a response.
func RemoveAlbumByID(ctx *gin.Context) {
	db := database.DBConn()
    id := ctx.Param("id")

	delete, err := db.Query(
		"DELETE FROM albums WHERE albumId=?", id)

	// if there is an error inserting, handle it
	if err != nil {
		fmt.Println("An error occured while attempting to delete album from db", err.Error())
	}

	defer delete.Close()
 	ctx.JSON(http.StatusOK, gin.H{"data": delete})
	defer db.Close()
}

// UpdateAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then removes that album as a response.
func UpdateAlbumByID(ctx *gin.Context) {
	db := database.DBConn()
	id := ctx.Param("id")
	var album entity.Album
	
	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	update, err := db.Query("UPDATE albums SET title=?,artist=?,price=? WHERE albumId=?",album.Title, album.Artist, album.Price, id)

	// if there is an error inserting, handle it
	if err != nil {
		fmt.Println("An error occured while attempting to update album on db", err.Error())
	}

	defer update.Close()
 	ctx.JSON(http.StatusOK, gin.H{"data": update})
	defer db.Close()
}