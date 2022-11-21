package repositories

import (
	"context"
	"database/sql"
	"e-commerce/models"
	"errors"
)

type AuthRepositoryInterface interface {
	Login(ctx context.Context, email string) (models.LoginDataResponse, error)
}

type AuthRepository struct {
	mysql *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		mysql: db,
	}
}

func (ar *AuthRepository) Login(ctx context.Context, email string) (models.LoginDataResponse, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = ?"

	var user models.LoginDataResponse
	err := ar.mysql.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.LoginDataResponse{}, errors.New("Data not found")
		}
		return models.LoginDataResponse{}, err
	}
	return user, err
}
