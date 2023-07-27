package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rahulsm20/go-crud-api/pkg/initializers"
	"github.com/rahulsm20/go-crud-api/pkg/middleware"
	"github.com/rahulsm20/go-crud-api/pkg/models"
)

func Signup(c *gin.Context) {

	//Validate Details
	err := middleware.ValidateUser(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}
	var body struct {
		Username string
		Password string
		Email    string
	}

	c.Bind(&body)
	//Create User
	user := models.User{Username: body.Username, Password: body.Password, Email: body.Email}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"Error": "Failed to create user",
		})
		return
	}

	c.JSON(201, gin.H{
		"User created": user,
	})

}

func DeleteUser(c *gin.Context) {
	var body struct {
		Username string
		Password string
		Email    string
	}

	c.Bind(&body)
	//Create User
	user := models.User{Username: body.Username, Password: body.Password, Email: body.Email}
	result := initializers.DB.Delete(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"Error": "Failed to delete user",
		})
		return
	}

	c.JSON(204, gin.H{
		"User deleted": user,
	})

}
