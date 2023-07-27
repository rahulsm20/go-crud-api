package middleware

import (
	"errors"
	"regexp"

	"github.com/gin-gonic/gin"
)

func ValidateUser(c *gin.Context) error {

	var user struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	c.Bind(&user)
	// var user UserDetails

	if user.Username == "" {
		return errors.New("please enter a valid username")
	}

	if user.Password == "" {
		return errors.New("please enter a valid password")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(emailRegex, user.Email); !matched {
		return errors.New("please enter a valid email")
	}

	return nil
}
