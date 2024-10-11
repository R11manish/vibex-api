package usecase

import (
	"fmt"
	models "vibex-api/internal/model"
)

type AuthUseCase interface {
	Login(signInRequest models.SignInRequest) (string, error)
	SignUp(signUpRequest models.SignUpRequest) error
}

type authUseCaseImpl struct {
}

func NewAuthUseCase() AuthUseCase {
	return &authUseCaseImpl{}
}

func (a *authUseCaseImpl) Login(signInRequest models.SignInRequest) (string, error) {
	fmt.Printf("SignInRequest: %+v\n", signInRequest.Identifier)

	return "some-token", nil
}

func (a *authUseCaseImpl) SignUp(signUpRequest models.SignUpRequest) error {

	fmt.Printf("SignUpRequest: %+v\n", signUpRequest)

	return nil
}
