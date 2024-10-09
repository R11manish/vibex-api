package routes

import (
	config "vibex-api/configs"
	"vibex-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.LoggerMiddleware(config.Logger))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
