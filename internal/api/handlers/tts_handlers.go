package handlers

import (
    "net/http"
    "github.com/akimisen/ai-task-manager/internal/models"
    "github.com/akimisen/ai-task-manager/internal/services"
    "github.com/gin-gonic/gin"
)

type TTSHandler struct {
    service *services.TTSService
}

func NewTTSHandler(service *services.TTSService) *TTSHandler {
    return &TTSHandler{service: service}
}

func (h *TTSHandler) CreateTask(c *gin.Context) {
    var task models.TTSTask
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

func (h *TTSHandler) GetTask(c *gin.Context) {
    taskID := c.Param("id")
    task, err := h.service.GetTask(taskID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, task)
}

func (h *TTSHandler) ListTasks(c *gin.Context) {
    tasks, err := h.service.ListTasks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, tasks)
}