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
	GetAllUsers(ctx context.Context) ([]models.UserResponse, error)
	DeleteUser(ctx context.Context, idUser int) error
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
	query := "SELECT id, username, email, password, gender, age, address, created_at, updated_at FROM users WHERE id = ?"

	err := ur.mysql.QueryRowContext(ctx, query, idUser).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Gender, &user.Age, &user.Address, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("Data not found")
		}
		return models.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) GetAllUsers(ctx context.Context) ([]models.UserResponse, error) {
	query := "SELECT id, username, email, gender, age, address, created_at, updated_at FROM users"

	rows, err := ur.mysql.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserResponse
	for rows.Next() {
		var user models.UserResponse
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Gender, &user.Age, &user.Address, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, idUser int) error {
	query := "DELETE FROM users WHERE id = ?"

	result, err := ur.mysql.ExecContext(ctx, query, idUser)
	if err != nil {
		return nil
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("data not found")
	}
	return nil
}
