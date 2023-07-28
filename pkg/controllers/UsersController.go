package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rahulsm20/go-crud-api/pkg/initializers"
	"github.com/rahulsm20/go-crud-api/pkg/middleware"
	"github.com/rahulsm20/go-crud-api/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	//Validate Details
	user, err := middleware.ValidateUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	result := initializers.DB.Create(user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
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
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to delete user",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"User deleted": user,
	})

}

func Signin(c *gin.Context) {

	var body struct {
		Username string
		Password string
		Email    string
	}

	c.Bind(&body)

	var user models.User
	initializers.DB.First(&user, "username=?", body.Username)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to find user",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Wrong password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})

}
