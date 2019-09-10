package controllers

import (
	b64 "encoding/base64"
	"github.com/ahmadkarlam-ralali/latihan_go/models"
	"github.com/ahmadkarlam-ralali/latihan_go/requests"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type AuthController struct {
	db *gorm.DB
}

func AuthControllerHandler(router *gin.Engine, db *gorm.DB) {
	handler := &AuthController{db: db}

	group := router.Group("/auth")
	{
		group.POST("/login", handler.Login)
	}
}

func (this *AuthController) Login(c *gin.Context) {
	var login requests.LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Wrong format",
		})
		return
	}

	user := this.authenticate(login, c)

	token := b64.StdEncoding.EncodeToString([]byte(user.Username))

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}

func (this *AuthController) authenticate(login requests.LoginRequest, c *gin.Context) models.User {
	var user models.User
	result := this.db.
		Where("username = ? and password = ?", login.Username, login.Password).
		First(&user)
	if result.Error != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Username or/and Password doesn't match",
		})
	}
	return user
}
