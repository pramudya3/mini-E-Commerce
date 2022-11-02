package repositories

import (
	"context"
	"database/sql"
	"e-commerce/models"
	"time"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, newUser models.User) error
	// GetUserById(ctx context.Context, userId int) (models.UserResponse, error)
	// GetAllUsers(ctx context.Context) models.UserResponse
	// DeleteUser(ctx context.Context, idToken int) error
	// UpdateUser(ctx context.Context, updateUser models.UserUpdate) (models.UserUpdateResponse, error)
}

type UserRepository struct {
	mysql *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		mysql: db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, newUser models.User) error {
	query := "INSERT INTO users (username, email, password, gender, age, address, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	_, err := ur.mysql.ExecContext(ctx, query, newUser.Username, newUser.Email, newUser.Password, newUser.Gender, newUser.Age, newUser.Address, now, now)
	if err != nil {
		return err
	}
	return nil
}
