package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowOrigins:    []string{"*"},
			AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:    []string{"Origin", "Authorization"},
			ExposeHeaders:   []string{"Content-Length"},
			MaxAge:          12 * time.Hour,
		})
	}
}
