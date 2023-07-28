package middleware

import (
	"errors"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/rahulsm20/go-crud-api/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func ValidateUser(c *gin.Context) (*models.User, error) {

	var user models.User

	c.Bind(&user)

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user.Password = string(hashedPassword)

	return &user, nil
}
