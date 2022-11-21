package services

import (
	"context"
	"e-commerce/models"
	"e-commerce/repositories"
	"errors"
	"time"
)

type CategoryServiceInterface interface {
	CreateCategory(ctx context.Context, newCategory models.NewCategory, idToken int) error
	GetCategoryById(ctx context.Context, id int) (models.Category, error)
	GetCategory(ctx context.Context) ([]models.Category, error)
	DeleteCategory(ctx context.Context, idToken int, id int) error
	UpdateCategory(ctx context.Context, updateCategory models.UpdateCategory, idToken, id int) (models.Category, error)
}

type CategoryService struct {
	categoryRepository repositories.CategoryRepositoryInterface
}

func NewCategoryService(categoryRepository repositories.CategoryRepositoryInterface) CategoryServiceInterface {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (cs *CategoryService) CreateCategory(ctx context.Context, newCategory models.NewCategory, idToken int) error {
	if newCategory.Name == "" {
		return errors.New("category name is required")
	}

	err := cs.categoryRepository.CreateCategory(ctx, newCategory, idToken)
	return err
}

func (cs *CategoryService) GetCategoryById(ctx context.Context, id int) (models.Category, error) {
	category, err := cs.categoryRepository.GetCategoryById(ctx, id)

	categoryResponse := models.Category{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
	return categoryResponse, err
}

func (cs *CategoryService) GetCategory(ctx context.Context) ([]models.Category, error) {
	category, err := cs.categoryRepository.GetCategory(ctx)
	return category, err
}

func (cs *CategoryService) DeleteCategory(ctx context.Context, idToken int, id int) error {
	err := cs.categoryRepository.DeleteCategory(ctx, idToken, id)
	return err
}

func (cs *CategoryService) UpdateCategory(ctx context.Context, updateCategory models.UpdateCategory, idToken, id int) (models.Category, error) {
	getCategory, err := cs.GetCategoryById(ctx, id)
	if err != nil {
		return models.Category{}, err
	}
	if updateCategory.Name != "" {
		getCategory.Name = updateCategory.Name
	}
	now := time.Now()
	getCategory.UpdatedAt = &now

	category, err := cs.categoryRepository.UpdateCategory(ctx, getCategory, idToken, id)
	responseUpdate := models.Category{
		Id:        getCategory.Id,
		Name:      category.Name,
		CreatedAt: getCategory.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
	return responseUpdate, err
}
