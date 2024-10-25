package middleware

import (
	"ai-task-manager/pkg/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := validation.ValidateRequest(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
