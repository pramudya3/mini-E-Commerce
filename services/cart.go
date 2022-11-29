package services

import (
	"context"
	"e-commerce/models"
	"e-commerce/repositories"
	"errors"
)

type CartServiceInterface interface {
	NewCart(ctx context.Context, newCart models.NewCart, idToken int) error
}

type CartService struct {
	cartRepo repositories.CartRepositoryInterface
}

func NewCartService(cartRepo repositories.CartRepositoryInterface) CartServiceInterface {
	return &CartService{
		cartRepo: cartRepo,
	}
}

func (cs *CartService) NewCart(ctx context.Context, newCart models.NewCart, idToken int) error {
	if newCart.ProductId == 0 {
		return errors.New("product_id is required")
	}
	if newCart.Quantity == 0 {
		return errors.New("quantity is required")
	}
	if newCart.SubTotal == "" {
		return errors.New("sub_total is required")
	}
	err := cs.cartRepo.NewCart(ctx, newCart, idToken)
	if err != nil {
		return err
	}
	return nil
}
