package http

import "github.com/gin-gonic/gin"

func NewRouter(taskHandler *TaskHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the TODO API!"})
	})
	tasks := r.Group("/tasks")
	{
		tasks.POST("", taskHandler.Create)
		tasks.PATCH("/:id/complete", taskHandler.Complete)
	}
	return r
}
