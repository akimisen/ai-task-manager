package middleware

import (
	"ai-task-manager/pkg/logger" // 修改为正确的 logger 包路径
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		clientIP := c.ClientIP()

		// 处理请求
		c.Next()

		// 计算请求处理时间
		latency := time.Since(start)
		statusCode := c.Writer.Status()

		// 记录请求信息
		log.Infof(
			"Method: %s | Path: %s | IP: %s | Status: %d | Latency: %v",
			method,
			path,
			clientIP,
			statusCode,
			latency,
		)

		// 如果有错误，记录错误信息
		if len(c.Errors) > 0 {
			log.Errorf("Request errors: %v", c.Errors)
		}
	}
}
