package repositories

import (
	"context"
	"database/sql"
	"e-commerce/models"
)

type CartRepositoryInterface interface {
	NewCart(ctx context.Context, newCart models.NewCart, idToken int) error
}

type CartRepository struct {
	mysql *sql.DB
}

func NewCartRepositor(db *sql.DB) *CartRepository {
	return &CartRepository{
		mysql: db,
	}
}

func (cr *CartRepository) NewCart(ctx context.Context, newCart models.NewCart, idToken int) error {
	query := "INSERT INTO carts (user_id, product_id, quantity, sub_total) VALUES(?, ?, ?, ?)"

	_, err := cr.mysql.ExecContext(ctx, query, idToken, newCart.ProductId, newCart.Quantity, newCart.SubTotal)
	if err != nil {
		return err
	}
	return nil
}
