package services

import (
	"context"
	"e-commerce/models"
	"e-commerce/repositories"
	"errors"
)

type OrderServiceInterface interface {
	NewOrder(ctx context.Context, newOrder models.NewOrder, idToken int) error
}

type OrderService struct {
	OrderRepository repositories.OrderRepositoryInterface
}

func NewOrderService(orderRepo repositories.OrderRepositoryInterface) OrderServiceInterface {
	return &OrderService{
		OrderRepository: orderRepo,
	}
}

func (os *OrderService) NewOrder(ctx context.Context, newOrder models.NewOrder, idToken int) error {
	if newOrder.CartId == 0 {
		return errors.New("cart_id is required")
	}
	if newOrder.PaymentMethod == "" {
		return errors.New("payment_method is required")
	}
	if newOrder.StatusOrder == "" {
		return errors.New("status_order is required")
	}
	if newOrder.TotalPrice == "" {
		return errors.New("total_price is required")
	}
	err := os.OrderRepository.NewOrder(ctx, newOrder, idToken)
	if err != nil {
		return err
	}
	return nil
}
