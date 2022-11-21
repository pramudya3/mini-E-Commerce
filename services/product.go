package services

import (
	"context"
	"e-commerce/models"
	"e-commerce/repositories"
	"errors"
)

type ProductServiceInterface interface {
	NewProduct(ctx context.Context, newProduct models.NewProduct, idToken int) error
	GetProductById(ctx context.Context, id int) (models.Product, error)
	GetAllProducts(ctx context.Context) ([]models.Product, error)
	DeleteProduct(ctx context.Context, idToken, id int) error
	UpdateProduct(ctx context.Context, updateProduct models.UpdateProduct, idToken, id int) (models.Category, error)
}

type ProductService struct {
	productRepository repositories.ProductRepositoryInterface
}

func NewProductService(productRepo repositories.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductService{
		productRepository: productRepo,
	}
}

func (ps *ProductService) NewProduct(ctx context.Context, newProduct models.NewProduct, idToken int) error {
	if newProduct.Name == "" {
		return errors.New("Name is required")
	}
	if newProduct.Price == "" {
		return errors.New("Price is required")
	}
	if newProduct.CategoryId == 0 {
		return errors.New("category_id is required")
	}
	if newProduct.Quantity == 0 {
		return errors.New("Quantity is required")
	}
	if newProduct.Status != "Available" {
		if newProduct.Status != "Not Available" {
			return errors.New("Status is required (Available or Not Available only)")
		}
	}
	err := ps.productRepository.NewProduct(ctx, newProduct, idToken)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductService) GetProductById(ctx context.Context, id int) (models.Product, error)
func (ps *ProductService) GetAllProducts(ctx context.Context) ([]models.Product, error)
func (ps *ProductService) DeleteProduct(ctx context.Context, idToken, id int) error
func (ps *ProductService) UpdateProduct(ctx context.Context, updateProduct models.UpdateProduct, idToken, id int) (models.Category, error)
