package repositories

import (
	"context"
	"database/sql"
	"e-commerce/models"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, newUser models.User) error
	GetUserById(ctx context.Context, userId int) (models.UserResponse, error)
	GetAllUsers(ctx context.Context) models.UserResponse
	DeleteUser(ctx context.Context, idToken int) error
	UpdateUser(ctx context.Context, updateUser models.UserUpdate) (models.UserUpdateResponse, error)
}

type UserRepository struct {
	mysql *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		mysql: db,
	}
}
