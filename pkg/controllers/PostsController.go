package controllers

import (
	"errors"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/rahulsm20/go-crud-api/pkg/initializers"
	"github.com/rahulsm20/go-crud-api/pkg/models"
	"gorm.io/gorm"
)

func CreatePost(c *gin.Context) {

	var body struct {
		Title   string
		Body    string
		User_id int
	}
	c.Bind(&body)

	var user models.User
	if body.User_id == 0 || body.Title == "" || body.Body == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Enter valid details",
		})
		return
	}
	err := initializers.DB.First(&user, "id=?", body.User_id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "User doesn't exist",
		})
		return
	}

	//Create Post
	post := models.Post{Title: body.Title, Body: body.Body, User_id: int(user.ID)}
	creationError := initializers.DB.Create(&post).Error

	if creationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": creationError,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"post": post,
	})
}

func FetchPostByID(c *gin.Context) {
	id := c.Param("id")
	match, err := regexp.MatchString(`\d+`, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Server Error",
		})
		// Log the error for debugging
		log.Println("Error matching ID:", err)
		return
	}

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Post ID",
		})
		return
	}
	var post models.Post
	initializers.DB.First(&post, id)
	if post.ID == 0 {
		c.JSON(400, gin.H{
			"Error": "Post doesn't exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func FetchAllPosts(c *gin.Context) {

	var post []models.Post
	initializers.DB.Find(&post)

	c.JSON(http.StatusOK, gin.H{
		"posts": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	match, _ := regexp.MatchString(`\d+`, id)

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Post ID",
		})
		return
	}

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	//Find post
	var post models.Post
	initializers.DB.First(&post, id)
	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid post ID",
		})
		return
	}

	//Update post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	match, _ := regexp.MatchString(`\d+`, id)

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Post ID",
		})
		return
	}
	var post models.Post

	if err := initializers.DB.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Respond with a message if the post doesn't exist
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Post not found",
			})
		} else {
			// Handle other database errors
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
		}
		return
	}

	if err := initializers.DB.Delete(&post, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Post ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}
