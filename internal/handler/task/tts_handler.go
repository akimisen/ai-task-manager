package task

import (
	"net/http"
	// "ai-task-manager/internal/models"
	"ai-task-manager/internal/service"

	"github.com/gin-gonic/gin"
)

type TTSHandler struct {
	ttsService *service.TTSService
}

func NewTTSHandler(ttsService *service.TTSService) *TTSHandler {
	return &TTSHandler{
		ttsService: ttsService,
	}
}

func (h *TTSHandler) CreateTask(c *gin.Context) {
	var request struct {
		Text string `json:"text" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.ttsService.CreateTask(request.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"task_id": task.ID, "status": task.Status})
}

func (h *TTSHandler) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := h.ttsService.GetTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task"})
		return
	}

	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No tasks available"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (h *TTSHandler) GetTaskStatus(c *gin.Context) {
	id := c.Param("id")
	task, err := h.ttsService.GetTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task"})
		return
	}

	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         task.ID,
		"status":     task.Status,
		"audio_data": task.AudioData,
	})
}

func (h *TTSHandler) ListTasks(c *gin.Context) {
	tasks, err := h.ttsService.ListTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list tasks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
