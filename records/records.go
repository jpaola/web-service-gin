package records

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	entity "github.com/jpaola/web-service-gin/entity"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(ctx *gin.Context) {
	// albums slice to seed record album data.
	albums := []entity.Album{}

	db, err := sql.Open("mysql", "<user>:<password>@/<databasename>")

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("DB Connection error", err.Error())
		return
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM albums")

	// if there is an error during query, handle it
	if err != nil {
		fmt.Println("Query Error", err.Error())
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
}

// addAlbum adds an album from JSON received in the request body.
func AddAlbum(ctx *gin.Context) {
	var newAlbum entity.Album

	db, err := sql.Open("mysql", "<user>:<password>@/<databasename>")

	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("DB Connection error", err.Error())
		return
	}

	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO albums (albumId,title,artist,price) VALUES (?,?,?,?)",
		newAlbum.ID, newAlbum.Title, newAlbum.Artist, newAlbum.Price)

	// if there is an error inserting, handle it
	if err != nil {
		fmt.Println("An error occured while attempting to insert new album to db", err.Error())
	}

	defer insert.Close()
 	ctx.JSON(http.StatusOK, gin.H{"data": insert})
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")
	
	// album to seed record album data.
	album := &entity.Album{}
	
	db, err := sql.Open("mysql", "<user>:<password>@/<databasename>")

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("DB Connection error", err.Error())
		return
	}

	defer db.Close()

	newAlbum, err := db.Query("SELECT * FROM albums where albumId=?", id)

	// if there is an error during query, handle it
	if err != nil {
		fmt.Println("Query Error", err.Error())
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
}

// removeAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then removes that album as a response.
func RemoveAlbumByID(ctx *gin.Context) {
    id := ctx.Param("id")

	db, err := sql.Open("mysql", "<user>:<password>@/<databasename>")

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("DB Connection error", err.Error())
		return
	}

	defer db.Close()

	delete, err := db.Query(
		"DELETE FROM albums WHERE albumId=?", id)

	// if there is an error inserting, handle it
	if err != nil {
		fmt.Println("An error occured while attempting to delete album from db", err.Error())
	}

	defer delete.Close()
 	ctx.JSON(http.StatusOK, gin.H{"data": delete})
}

// updateAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then removes that album as a response.
func UpdateAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id") 
	var album entity.Album
	
	db, err := sql.Open("mysql", "<user>:<password>@/<databasename>")
	
	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("DB Connection error", err.Error())
		return
	}
	
	defer db.Close()
	
	update, err := db.Query("UPDATE albums SET title=?,artist=?,price=? WHERE albumId=?",album.Title, album.Artist, album.Price, id)

	// if there is an error inserting, handle it
	if err != nil {
		fmt.Println("An error occured while attempting to update album on db", err.Error())
	}

	defer update.Close()
 	ctx.JSON(http.StatusOK, gin.H{"data": update})
}