package services

import (
	"context"
	"e-commerce/helpers"
	"e-commerce/models"
	"e-commerce/repositories"
	"errors"
	"time"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, newUser models.CreateUserRequest) error
	GetUserById(ctx context.Context, idUser int) (models.UserResponse, error)
	GetAllUsers(ctx context.Context) ([]models.UserResponse, error)
	DeleteUser(ctx context.Context, idToken int) error
	UpdateUser(ctx context.Context, updateUser models.UserUpdateRequest, idToken int) (models.UserResponse, error)
}

type UserService struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserService(userRepo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepository: userRepo,
	}
}

func (us *UserService) CreateUser(ctx context.Context, newUser models.CreateUserRequest) error {
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
		Id:        user.Id,
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

func (us *UserService) DeleteUser(ctx context.Context, idToken int) error {
	err := us.userRepository.DeleteUser(ctx, idToken)
	return err
}

func (us *UserService) UpdateUser(ctx context.Context, updateUser models.UserUpdateRequest, idToken int) (models.UserResponse, error) {
	getUser, err := us.userRepository.GetUserById(ctx, idToken)
	if err != nil {
		return models.UserResponse{}, err
	}
	if updateUser.Username != "" {
		getUser.Username = updateUser.Username
	}
	if updateUser.Email != "" {
		getUser.Email = updateUser.Email
	}
	if updateUser.Password != "" {
		getUser.Password = updateUser.Password
	}
	if updateUser.Gender != "" {
		if updateUser.Gender != "Male" {
			if updateUser.Gender != "Female" {
				return models.UserResponse{}, errors.New("Update gender is only Male and Female")
			}
		}
		getUser.Gender = updateUser.Gender
	}
	if updateUser.Age != 0 {
		getUser.Age = updateUser.Age
	}
	if updateUser.Address != "" {
		getUser.Address = updateUser.Address
	}

	layoutFormat := "2006-01-02T15:04:05"
	value := time.Now().Local().Format(layoutFormat)
	now, _ := time.Parse(layoutFormat, value)
	getUser.UpdatedAt = &now

	user, err := us.userRepository.UpdateUser(ctx, getUser, idToken)
	responseUpdate := models.UserResponse{
		Id:        getUser.Id,
		Username:  user.Username,
		Email:     user.Email,
		Gender:    user.Gender,
		Age:       user.Age,
		Address:   user.Address,
		CreatedAt: getUser.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return responseUpdate, err
}
