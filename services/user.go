package services

import (
	"context"
	"e-commerce/models"
	"e-commerce/repositories"
	"errors"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, newUser models.User) error
}

type UserService struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserService(userRepo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepository: userRepo,
	}
}

func (us *UserService) CreateUser(ctx context.Context, newUser models.User) error {
	if newUser.Username == "" {
		return errors.New("Username is required")
	}
	if newUser.Email == "" {
		return errors.New("Email is required")
	}
	if newUser.Password == "" {
		return errors.New("Password is required")
	}
	if newUser.Gender != "Male" {
		if newUser.Gender != "Female" {
			return errors.New("Gender is only Male and Female")
		}
	}
	if newUser.Age == 0 {
		return errors.New("Age is required")
	}
	if newUser.Address == "" {
		return errors.New("Address is required")
	}

	err := us.userRepository.CreateUser(ctx, newUser)
	return err
}
