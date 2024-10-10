package routes

import (
	config "vibex-api/configs"
	"vibex-api/internal/controller"
	"vibex-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.LoggerMiddleware(config.Logger))

	authController := controller.NewAuthController()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// auth
	auth := router.Group("/auth")
	{
		auth.POST("/signup", authController.SignUpHandler)
		auth.GET("/login", authController.LoginHandler)
	}

	return router
}
