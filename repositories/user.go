package repositories

import (
	"context"
	"database/sql"
	"e-commerce/models"
	"errors"
	"time"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, newUser models.User) error
	GetUserById(ctx context.Context, idUser int) (models.User, error)
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
	query := "INSERT INTO users(username, email, password, gender, age, address, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	_, err := ur.mysql.ExecContext(ctx, query, newUser.Username, newUser.Email, newUser.Password, newUser.Gender, newUser.Age, newUser.Address, now, now)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserById(ctx context.Context, idUser int) (models.User, error) {
	var user models.User
	query := "SELECT username, email, password, gender, age, address, created_at, updated_at FROM users WHERE id = ?"

	err := ur.mysql.QueryRowContext(ctx, query, idUser).Scan(&user.Username, &user.Email, &user.Password, &user.Gender, &user.Age, &user.Address, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("Data not found")
		}
		return models.User{}, err
	}
	return user, nil
}
