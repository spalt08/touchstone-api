package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Setup will add middlewares to gin instance. It used both for server running and testing
func Setup(handler *gin.Engine) {
	handler.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET"},
		AllowHeaders:    []string{"Origin", "Content-Type"},
	}))
}
