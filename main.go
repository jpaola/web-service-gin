package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// album represents data about a record album.
type Album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// Database credentials
var dbuser = "nameOfUser"
var dbpass = "somePasscode" 
var dbname = "nameOfDB"


// getAlbums responds with the list of all albums as JSON.
func GetAlbums(ctx *gin.Context) {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

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
	
	// albums slice to seed record album data.
 	albums := []Album{}

	for results.Next() {
		var a Album
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
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

    var newAlbum Album

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
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	id := ctx.Param("id")

	// album to seed record album data.
 	album := &Album{}

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

	fmt.Println(album)
	ctx.IndentedJSON(http.StatusOK, album)
}

// removeAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then removes that album as a response.
func RemoveAlbumByID(ctx *gin.Context) {
    id := ctx.Param("id")

	// albums slice to seed record album data.
 	albums := []Album{}

	// Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for i, album := range albums {
        if album.ID == id {
			copy(albums[i:], albums[i+1:]) // Shift a[i+1:] left one index.
			// albums[len(albums)-1] = {""}     // Erase last element (write zero value).
			albums = albums[:len(albums)-1]     // Truncate slice.
            ctx.IndentedJSON(http.StatusOK, album)
            return
        }
    }
    ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album could not be deleted"})
}

// updateAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then removes that album as a response.
func UpdateAlbumByID(ctx *gin.Context) {

    id := ctx.Param("id")
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

    var album Album

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("DB Connection error", err.Error())
		return
	}

	defer db.Close()

	update, err := db.Query(
		"UPDATE albums SET title=?,artist=?,price=? WHERE albumId=?",
		album.Title, album.Artist, album.Price, id)

	// if there is an error inserting, handle it
	if err != nil {
		fmt.Println("An error occured while attempting to update album on db", err.Error())
	}

	defer update.Close()
 	ctx.JSON(http.StatusOK, gin.H{"data": update})
}

func main() {

	// initialize gin router using Default
    router := gin.Default()

	router.GET("/albums", GetAlbums)
    router.GET("/albums/:id", GetAlbumByID)
    router.POST("/albums", AddAlbum)
    router.DELETE("/albums/:id", RemoveAlbumByID)
    router.PUT("/albums/:id", UpdateAlbumByID)

    router.Run("localhost:8080")
}