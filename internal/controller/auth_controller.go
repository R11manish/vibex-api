package controller

import "github.com/gin-gonic/gin"

type AuthController interface {
	LoginHandler(c *gin.Context)
	SignUpHandler(c *gin.Context)
}

type authControllerIml struct {
}

func NewAuthController() AuthController {
	return &authControllerIml{}
}

func (a *authControllerIml) LoginHandler(c *gin.Context) {

}

func (a *authControllerIml) SignUpHandler(c *gin.Context) {

}
