package services

import (
	"context"
	"e-commerce/helpers"
	"e-commerce/models"
	"e-commerce/repositories"
	"errors"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, newUser models.User) error
	GetUserById(ctx context.Context, idUser int) (models.UserResponse, error)
	GetAllUsers(ctx context.Context) ([]models.UserResponse, error)
	DeleteUser(ctx context.Context, idUser int) error
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
	if newUser.Password == "" {
		return errors.New("Password is required")
	}
	//hashing password
	pass, errHash := helpers.HashPassword(newUser.Password)
	if errHash != nil {
		return errors.New("failed")
	}
	newUser.Password = pass

	err := us.userRepository.CreateUser(ctx, newUser)
	return err
}

func (us *UserService) GetUserById(ctx context.Context, idUser int) (models.UserResponse, error) {
	user, err := us.userRepository.GetUserById(ctx, idUser)

	userResponse := models.UserResponse{
		Username:  user.Username,
		Email:     user.Email,
		Gender:    user.Gender,
		Age:       user.Age,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return userResponse, err
}

func (us *UserService) GetAllUsers(ctx context.Context) ([]models.UserResponse, error) {
	user, err := us.userRepository.GetAllUsers(ctx)
	return user, err
}

func (us *UserService) DeleteUser(ctx context.Context, idUser int) error {
	err := us.userRepository.DeleteUser(ctx, idUser)
	return err
}
