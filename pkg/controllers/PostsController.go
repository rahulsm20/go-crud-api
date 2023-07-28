package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulsm20/go-crud-api/pkg/initializers"
	"github.com/rahulsm20/go-crud-api/pkg/models"
)

func CreatePost(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	//Create Post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"post": post,
	})
}

func FetchPostByID(c *gin.Context) {

	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)
	if post.ID == 0 {
		c.JSON(400, gin.H{
			"Error": "Invalid post ID",
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

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Post deleted successfully",
	})
}
