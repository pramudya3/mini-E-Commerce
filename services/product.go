package services

import (
	"context"
	"e-commerce/models"
	"e-commerce/repositories"
	"errors"
	"strconv"
	"time"
)

type ProductServiceInterface interface {
	NewProduct(ctx context.Context, newProduct models.NewProduct, idToken int) error
	GetProductById(ctx context.Context, id int) (models.ResponseProduct, error)
	GetAllProducts(ctx context.Context) ([]models.ResponseProduct, error)
	DeleteProduct(ctx context.Context, idToken, id int) error
	UpdateProduct(ctx context.Context, updateProduct models.UpdateProduct, idToken, id int) (models.ResponseProduct, error)
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

func (ps *ProductService) GetProductById(ctx context.Context, id int) (models.ResponseProduct, error) {
	product, err := ps.productRepository.GetProductById(ctx, id)

	productResponse := models.ResponseProduct{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		CategoryId:  product.CategoryId,
		Quantity:    product.Quantity,
		Status:      product.Status,
		Description: product.Description,
	}
	return productResponse, err
}

func (ps *ProductService) GetAllProducts(ctx context.Context) ([]models.ResponseProduct, error) {
	product, err := ps.productRepository.GetAllProducts(ctx)
	return product, err
}

func (ps *ProductService) DeleteProduct(ctx context.Context, idToken, id int) error {
	err := ps.productRepository.DeleteProduct(ctx, idToken, id)
	return err
}

func (ps *ProductService) UpdateProduct(ctx context.Context, updateProduct models.UpdateProduct, idToken, id int) (models.ResponseProduct, error) {
	getProduct, errGet := ps.productRepository.GetProductById(ctx, id)
	idProduct := strconv.Itoa(updateProduct.CategoryId)

	if errGet != nil {
		return models.ResponseProduct{}, errGet
	}
	if updateProduct.Name != "" {
		getProduct.Name = updateProduct.Name
	}
	if updateProduct.Price != "" {
		getProduct.Price = updateProduct.Price
	}
	if updateProduct.CategoryId != 0 {
		getProduct.CategoryId = idProduct
	}
	if updateProduct.Quantity != 0 {
		getProduct.Quantity = updateProduct.Quantity
	}
	if updateProduct.Status != "" {
		if updateProduct.Status != "Available" {
			if updateProduct.Status != "Not Available" {
				return models.ResponseProduct{}, errors.New("update status only Available and Not Available")
			}
		}
		getProduct.Status = updateProduct.Status
	}
	if updateProduct.Description != "" {
		getProduct.Description = updateProduct.Description
	}

	layoutFormat := "2006-01-02T15:04:05"
	value := time.Now().Local().Format(layoutFormat)
	now, _ := time.Parse(layoutFormat, value)
	getProduct.UpdatedAt = &now

	product, err := ps.productRepository.UpdateProduct(ctx, getProduct, idToken, id)

	responseUpdate := models.ResponseProduct{
		Id:          getProduct.Id,
		Name:        product.Name,
		Price:       product.Price,
		CategoryId:  product.CategoryId,
		Quantity:    product.Quantity,
		Status:      product.Status,
		Description: product.Description,
	}
	return responseUpdate, err
}
