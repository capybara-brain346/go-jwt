package controllers

import (
	"net/http"

	"github.com/capybara-brain346/go-jwt/initializers"
	"github.com/capybara-brain346/go-jwt/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})

}

func Login(c *gin.Context) {}
func GetUser(c *gin.Context) {}
func UpdateUser(c *gin.Context) {}
func DeleteUser(c *gin.Context) {}
func Logout(c *gin.Context) {}
func RequireAuth(c *gin.Context) {}
