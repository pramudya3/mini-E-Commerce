package repositories

import (
	"context"
	"database/sql"
	"e-commerce/models"
)

type OrderRepositoryInterface interface {
	NewOrder(ctx context.Context, newOrder models.NewOrder, idToken int) error
}

type OrderRepository struct {
	mysql *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		mysql: db,
	}
}

func (or *OrderRepository) NewOrder(ctx context.Context, newOrder models.NewOrder, idToken int) error {
	query := "INSERT INTO orders (cart_id, payment_method, status_order, total_price) VALUES (?, ?, ?, ?)"

	_, err := or.mysql.ExecContext(ctx, query, idToken, newOrder.PaymentMethod, newOrder.StatusOrder, newOrder.TotalPrice)
	if err != nil {
		return err
	}
	return nil
}
