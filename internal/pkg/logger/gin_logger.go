package logger

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GinLoggerMiddleware(logger *Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		status := ctx.Writer.Status()

		fields := logrus.Fields{
			"method":      ctx.Request.Method,
			"url":         ctx.Request.URL.String(),
			"status":      status,
			"latency":     time.Since(start),
			"clientIP":    ctx.ClientIP(),
			"userAgent":   ctx.Request.UserAgent(),
			"requestSize": ctx.Request.ContentLength,
		}

		switch {
		case status >= http.StatusInternalServerError:
			logger.WithFields(fields).Error("Request resulted in a server error")
		case status >= http.StatusBadRequest && status < http.StatusInternalServerError:
			logger.WithFields(fields).Warn("Request resulted in a client error")
		default:
			logger.WithFields(fields).Info("Request ended successfully")
		}
	}
}
