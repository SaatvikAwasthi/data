package requestLogger

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := context.Request.Method
		reqUri := context.Request.URL.Path
		queryParams := context.Request.URL.RawQuery
		userAgent := context.Request.UserAgent()
		statusCode := context.Writer.Status()
		clientIP := context.ClientIP()
		format := "| status:%3d | method:%s | path:%s | query:%s | ip:%s | user-agent:%s | latency:%13v | context-errors:%v\n"
		log.Printf(format, statusCode, reqMethod, reqUri, queryParams, clientIP, userAgent, latencyTime, context.Errors)
	}
}
