package routes

import (
	config "vibex-api/configs"
	"vibex-api/internal/controller"
	"vibex-api/internal/middleware"
	"vibex-api/internal/repository"
	"vibex-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.LoggerMiddleware(config.Logger))
	db := config.GetDB()

	//repo initilization

	userRepo := repository.NewUserRepository(db)
	// auth
	authUseCase := usecase.NewAuthUseCase(userRepo)
	authController := controller.NewAuthController(authUseCase)

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
