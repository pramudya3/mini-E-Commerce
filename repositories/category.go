package repositories

import (
	"context"
	"database/sql"
	"e-commerce/models"
	"errors"
)

type CategoryRepositoryInterface interface {
	CreateCategory(ctx context.Context, newCategory models.NewCategory, idToken int) error
	GetCategoryById(ctx context.Context, id int) (models.Category, error)
	GetCategory(ctx context.Context) ([]models.Category, error)
	DeleteCategory(ctx context.Context, idToken int, id int) error
	UpdateCategory(ctx context.Context, updateCategory models.Category, idToken, id int) (models.Category, error)
}

type CategoryRepository struct {
	mysql *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		mysql: db,
	}
}

func (cr *CategoryRepository) CreateCategory(ctx context.Context, newCategory models.NewCategory, idToken int) error {
	query := "INSERT INTO categories (name, user_id) VALUES (?, ?)"

	_, err := cr.mysql.ExecContext(ctx, query, newCategory.Name, idToken)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CategoryRepository) GetCategoryById(ctx context.Context, id int) (models.Category, error) {
	var category models.Category
	query := "SELECT id, name, created_at FROM categories WHERE id = ?"

	err := cr.mysql.QueryRowContext(ctx, query, id).Scan(&category.Id, &category.Name, &category.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Category{}, errors.New("data not found")
		}
		return models.Category{}, err
	}
	return category, nil
}

func (cr *CategoryRepository) GetCategory(ctx context.Context) ([]models.Category, error) {
	query := "SELECT id, name, created_at, updated_at FROM categories"

	rows, err := cr.mysql.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (cr *CategoryRepository) DeleteCategory(ctx context.Context, idToken int, id int) error {
	query := "DELETE FROM categories WHERE user_id = ? AND id = ?"

	result, err := cr.mysql.ExecContext(ctx, query, idToken, id)
	if err != nil {
		return nil
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (cr *CategoryRepository) UpdateCategory(ctx context.Context, updateCategory models.Category, idToken, id int) (models.Category, error) {
	query := "UPDATE categories SET name = ?, updated_at = ? WHERE user_id = ? AND id = ?"

	result, err := cr.mysql.ExecContext(ctx, query, updateCategory.Name, updateCategory.UpdatedAt, idToken, id)
	if err != nil {
		return models.Category{}, err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return models.Category{}, errors.New("category not found")
	}
	return updateCategory, nil
}
