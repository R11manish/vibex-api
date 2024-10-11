package controller

import (
	"net/http"
	models "vibex-api/internal/model"
	"vibex-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	LoginHandler(c *gin.Context)
	SignUpHandler(c *gin.Context)
}

type authControllerIml struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthController(authUseCase usecase.AuthUseCase) AuthController {
	return &authControllerIml{authUseCase: authUseCase}
}

func (a *authControllerIml) LoginHandler(c *gin.Context) {
	var signInRequest models.SignInRequest

	if err := c.ShouldBindJSON(&signInRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := a.authUseCase.Login(signInRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *authControllerIml) SignUpHandler(c *gin.Context) {

}
