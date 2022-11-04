package services

import (
	"context"
	"e-commerce/helpers"
	"e-commerce/middlewares"
	"e-commerce/models"
	"e-commerce/repositories"
	"errors"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, userLogin models.LoginRequest) (models.LoginResponse, error)
}

type AuthService struct {
	authRepository repositories.AuthRepositoryInterface
}

func NewAuthService(authRepo repositories.AuthRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		authRepository: authRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, userLogin models.LoginRequest) (models.LoginResponse, error) {
	if userLogin.Email == "" {
		return models.LoginResponse{}, errors.New("Email is required")
	}
	user, err := as.authRepository.Login(ctx, userLogin.Email)
	if err != nil {
		return models.LoginResponse{}, err
	}

	if !helpers.CheckPasswordHash(userLogin.Password, user.Password) {
		return models.LoginResponse{}, errors.New("Password incorrect")
	}

	token, err := middlewares.CreateToken(user.Id)
	if err != nil {
		return models.LoginResponse{}, err
	}

	loginResponse := models.LoginResponse{
		Id:       user.Id,
		Username: user.Username,
		Token:    token,
	}
	return loginResponse, err
}
