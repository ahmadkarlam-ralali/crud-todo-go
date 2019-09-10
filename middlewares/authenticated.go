package middlewares

import (
	b64 "encoding/base64"
	"github.com/ahmadkarlam-ralali/latihan_go/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

func Authenticate(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")
		username, _ := b64.StdEncoding.DecodeString(token)

		var user models.User
		result := db.First(&user, "username = ?", string(username))
		if result.Error != nil {
			c.Abort()
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "Invalid token",
			})
			return
		}

		c.Set("user_id", user.ID)
		c.Next()
	}
}
