package pkg

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/BAITOEYSRN/test-Technical-Skill/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		log.Printf("[%s] %s | %d | %v", method, path, statusCode, latency)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[PANIC] %v\n%s", r, debug.Stack())
				response.ResponseJsonWithCode(
					c,
					http.StatusInternalServerError,
					uuid.New(),
					"error",
					"Internal Server Error",
					nil,
				)

				c.Abort()
			}
		}()

		c.Next()
	}
}
