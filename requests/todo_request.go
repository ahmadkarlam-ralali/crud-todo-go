package requests

type StoreTodoRequest struct {
	Title string `form:"title" json:"title" binding:"required"`
}

type UpdateTodoRequest struct {
	Title string `form:"title" json:"title"`
	Status int `form:"status" json:"status" binding:"required"`
}