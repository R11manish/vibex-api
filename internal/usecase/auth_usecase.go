package usecase

import (
	"errors"
	"fmt"
	models "vibex-api/internal/model"
	"vibex-api/internal/repository"
	utils "vibex-api/internal/utlis"
)

type AuthUseCase interface {
	Login(signInRequest models.SignInRequest) (string, error)
	SignUp(signUpRequest models.SignUpRequest) error
}

type authUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewAuthUseCase(userRepository repository.UserRepository) AuthUseCase {
	return &authUseCaseImpl{
		userRepository: userRepository,
	}
}

func (a *authUseCaseImpl) Login(signInRequest models.SignInRequest) (string, error) {
	fmt.Printf("SignInRequest: %+v\n", signInRequest.Identifier)

	return "some-token", nil
}

func (a *authUseCaseImpl) SignUp(signUpRequest models.SignUpRequest) error {

	// Check if the user already exists
	existingUser, _ := a.userRepository.FindUserByUsername(signUpRequest.Username)
	if existingUser != nil {
		return errors.New("username already taken")
	}

	// Check if the email is already registered
	existingUser, _ = a.userRepository.FindUserByEmail(signUpRequest.Email)
	if existingUser != nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := utils.HashPassword(signUpRequest.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	newUser := &models.User{
		ID:       2232323232,
		Username: signUpRequest.Username,
		Name:     deref(signUpRequest.Name),
		Email:    signUpRequest.Email,
		StatusID: 1,
		Password: hashedPassword,
	}

	if err := a.userRepository.CreateUser(newUser); err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func deref(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
