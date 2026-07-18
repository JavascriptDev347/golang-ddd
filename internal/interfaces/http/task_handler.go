package http

import (
	"net/http"

	appTask "github.com/JavascriptDev347/golang-ddd/internal/application/task"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskHandler struct {
	createUC   *appTask.CreateTaskUseCase
	completeUC *appTask.CompleteTaskUseCase
}

func NewTaskHandler(createUC *appTask.CreateTaskUseCase, completeUC *appTask.CompleteTaskUseCase) *TaskHandler {
	return &TaskHandler{createUC: createUC, completeUC: completeUC}
}

type createTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func (h *TaskHandler) Create(c *gin.Context) {
	var req createTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t, err := h.createUC.Execute(c.Request.Context(), appTask.CreateTaskInput{
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, t)
}
func (h *TaskHandler) Complete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.completeUC.Execute(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
