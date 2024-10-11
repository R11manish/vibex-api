package usecase

import (
	"errors"
	"fmt"
	models "vibex-api/internal/model"
	"vibex-api/internal/repository"
	"vibex-api/internal/services"
	utils "vibex-api/internal/utlis"
)

type AuthUseCase interface {
	Login(signInRequest models.SignInRequest) (string, error)
	SignUp(signUpRequest models.SignUpRequest) error
}

type authUseCaseImpl struct {
	userRepository repository.UserRepository
	jwtService     services.JWTService
}

func NewAuthUseCase(userRepository repository.UserRepository, jwtService services.JWTService) AuthUseCase {
	return &authUseCaseImpl{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

func (a *authUseCaseImpl) Login(signInRequest models.SignInRequest) (string, error) {

	var user *models.User
	var err error

	if utils.IsEmail(signInRequest.Identifier) {
		user, err = a.userRepository.FindUserByEmail(signInRequest.Identifier)
	} else {
		user, err = a.userRepository.FindUserByUsername(signInRequest.Identifier)
	}

	if err != nil {
		return "", fmt.Errorf("user not found: %v", err)
	}
	if user == nil {
		return "", errors.New("invalid email or password")
	}

	err = utils.CheckPassword(user.Password, signInRequest.Password)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := a.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
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

	// Check password validation
	if err := utils.CheckPasswordStrength(signUpRequest.Password); err != nil {
		return fmt.Errorf("password validation failed: %v", err)
	}

	hashedPassword, err := utils.HashPassword(signUpRequest.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	id, err := utils.GenerateID()
	if err != nil {
		return fmt.Errorf("failed to generate id: %v", err)
	}

	newUser := &models.User{
		ID:       id,
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
