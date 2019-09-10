package controllers

import (
	"github.com/ahmadkarlam-ralali/latihan_go/middlewares"
	"github.com/ahmadkarlam-ralali/latihan_go/models"
	"github.com/ahmadkarlam-ralali/latihan_go/requests"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type TodosController struct {
	db *gorm.DB
}

func TodosControllerHandler(router *gin.Engine, db *gorm.DB) {
	handler := &TodosController{db: db}

	group := router.Group("/todos")

	group.Use(middlewares.Authenticate(db))
	{
		group.GET("/", handler.GetAll)
		group.POST("/", handler.Store)
		group.PUT("/:id", handler.Update)
		group.DELETE("/:id", handler.Destroy)
	}
}

func (this *TodosController) GetAll(c *gin.Context) {
	userId := c.MustGet("user_id").(uint)
	var todos []models.Todo
	this.db.Find(&todos, "user_id = ?", userId)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   todos,
	})
}

func (this *TodosController) Store(c *gin.Context) {
	var request requests.StoreTodoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Wrong format",
		})
		return
	}

	todo := models.Todo{
		Title:  request.Title,
		Status: 0,
		UserId: c.MustGet("user_id").(uint),
	}
	this.db.Create(&todo)

	c.JSON(http.StatusOK, gin.H{
		"data":    todo,
		"status":  "success",
		"message": "Todo created",
	})
}

func (this *TodosController) Update(c *gin.Context) {
	var request requests.UpdateTodoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Wrong format",
		})
		return
	}

	var todo models.Todo
	id := c.Param("id")
	userId := c.MustGet("user_id").(uint)

	result := this.db.First(&todo, "id = ? and user_id = ?", id, userId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "To Do not found",
		})
		return
	}

	if request.Title != "" {
		todo.Title = request.Title
	}
	todo.Status = request.Status

	this.db.Model(&todo).Updates(todo)

	c.JSON(http.StatusOK, gin.H{
		"data":    todo,
		"status":  "success",
		"message": "Todo updated",
	})
}

func (this *TodosController) Destroy(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	userId := c.MustGet("user_id").(uint)

	result := this.db.First(&todo, "id = ? and user_id = ?", id, userId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "To Do not found",
		})
		return
	}

	this.db.Delete(todo)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "To Do deleted",
	})
}
