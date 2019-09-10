package routes

import (
	"github.com/ahmadkarlam-ralali/latihan_go/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	controllers.AuthControllerHandler(r, db)

	controllers.TodosControllerHandler(r, db)

	return r
}
