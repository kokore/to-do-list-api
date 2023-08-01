package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hugeman/todolist/internal/config"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		if origin := c.Request.Header.Get("Origin"); config.Config.App.Env == "local" {
			c.Header("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
