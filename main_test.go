// Golang REST API unit testing program
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jpaola/web-service-gin/records"
)

func TestGetAlbums(t *testing.T) {
    // Switch to test mode so you don't get such noisy output
    gin.SetMode(gin.TestMode)

    // Setup your router, just like you did in your main function, and
    // register your routes
    router := gin.Default()
    router.GET("/albums", records.GetAlbums)

    // Create the mock request you'd like to test. Make sure the second argument
    // here is the same as one of the routes you defined in the router setup
    // block!
    req, err := http.NewRequest(http.MethodGet, "/albums", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    // Create a response recorder so you can inspect the response
    res := httptest.NewRecorder()

    // Perform the request
    router.ServeHTTP(res, req)
    fmt.Println(res.Body)

    // Check to see if the response was what you expected
    if res.Code == http.StatusOK {
        t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, res.Code)
    } else {
        t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, res.Code)
    }
}
func TestAddAlbum(t *testing.T) {
    // Switch to test mode so you don't get such noisy output
    gin.SetMode(gin.TestMode)

    // Setup your router, just like you did in your main function, and
    // register your routes
    router := gin.Default()
    router.POST("/albums", records.AddAlbum)

    // Create the mock request you'd like to test. Make sure the second argument
    // here is the same as one of the routes you defined in the router setup
    // block!
    req, err := http.NewRequest(http.MethodPost, "/albums", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    // Create a response recorder so you can inspect the response
    res := httptest.NewRecorder()

    // Perform the request
    router.ServeHTTP(res, req)
    fmt.Println(res.Body)

    // Check to see if the response was what you expected
    if res.Code == http.StatusOK {
        t.Logf("Expected to get status %d to be same as %d\n", http.StatusOK, res.Code)
    } else {
        t.Fatalf("Expected to get status %d, but instead got %d\n", http.StatusOK, res.Code)
    }
}
