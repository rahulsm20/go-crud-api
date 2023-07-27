package middleware

import (
	"errors"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/rahulsm20/go-crud-api/pkg/models"
)

func ValidateUser(c *gin.Context) (*models.User, error) {

	var user models.User

	c.Bind(&user)
	// var user UserDetails

	if user.Username == "" {
		return nil, errors.New("please enter a valid username")
	}

	if user.Password == "" {
		return nil, errors.New("please enter a valid password")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(emailRegex, user.Email); !matched {
		return nil, errors.New("please enter a valid email")
	}

	return &user, nil
}
