package handlers

import (
	"net/http"

	"github.com/akimisen/ai-task-manager/internal/models"
	"github.com/akimisen/ai-task-manager/internal/services"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	service *services.VideoService
}

func NewVideoHandler(service *services.VideoService) *VideoHandler {
	return &VideoHandler{service: service}
}

func (h *VideoHandler) CreateTask(c *gin.Context) {
	var task models.VideoTask
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *VideoHandler) GetTask(c *gin.Context) {
	taskID := c.Param("id")
	task, err := h.service.GetTask(taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *VideoHandler) ListTasks(c *gin.Context) {
	tasks, err := h.service.ListTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
