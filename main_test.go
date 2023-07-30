package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rahulsm20/go-crud-api/pkg/controllers"
)

// Test for CreatePost function
func TestCreatePost(t *testing.T) {
	// Create a new Gin router
	router := gin.New()

	// Define a sample post payload for the test
	postPayload := map[string]string{
		"Title": "Test Title",
		"Body":  "Test Body",
	}

	// Convert postPayload to JSON format
	payloadBytes, _ := json.Marshal(postPayload)

	// Create a request with the JSON payload
	req, err := http.NewRequest("POST", "/posts/", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Perform the request to the handler
	router.POST("/posts/", controllers.CreatePost)
	router.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %v but got %v", http.StatusCreated, status)
	}

	// Check the response body to ensure it contains the expected post details
	// expectedResponse := `{
	// 	"post": {
	// 		"ID": 0,
	// 		"CreatedAt": "2006-01-02T15:04:05.000000000Z",
	// 		"UpdatedAt": "2006-01-02T15:04:05.000000000Z",
	// 		"DeletedAt": null,
	// 		"Title": "Test Title",
	// 		"Body": "Test Body"}
	// }`
	// if rr.Body.String() != expectedResponse {
	// 	t.Errorf("Expected response body: %v\nBut got: %v", expectedResponse, rr.Body.String())
	// }
}

func TestGetAllPost(t *testing.T) {
	router := gin.New()

	// Create a request with the JSON payload
	req, err := http.NewRequest("GET", "/posts/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Set the request Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	cookie := http.Cookie{Name: "Authorization", Value: os.Getenv("COOKIE")}
	req.AddCookie(&cookie)
	// Perform the request to the handler
	router.GET("/posts/", controllers.FetchAllPosts)
	router.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v but got %v", http.StatusOK, status)
	}

}

func TestGetPostByID(t *testing.T) {
	router := gin.New()

	req, err := http.NewRequest("GET", "/posts/9", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	cookie := http.Cookie{Name: "Authorization", Value: os.Getenv("COOKIE")}
	req.AddCookie(&cookie)

	// Perform the request to the handler
	router.GET("/posts/:id", controllers.FetchPostByID)
	router.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v but got %v", http.StatusOK, status)
	}

}
