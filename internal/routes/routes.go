package routes

import (
	"os"
	config "vibex-api/configs"
	"vibex-api/internal/controller"
	"vibex-api/internal/middleware"
	"vibex-api/internal/repository"
	"vibex-api/internal/services"
	"vibex-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.LoggerMiddleware(config.Logger))
	db := config.GetDB()
	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	//service initilization

	jwtService := services.NewJWTService(jwtSecret, "vibex-api")

	//repo initilization
	userRepo := repository.NewUserRepository(db)
	// auth
	authUseCase := usecase.NewAuthUseCase(userRepo, jwtService)
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
