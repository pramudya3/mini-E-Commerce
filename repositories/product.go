package repositories

import (
	"context"
	"database/sql"
	"e-commerce/models"
)

type ProductRepositoryInterface interface {
	NewProduct(ctx context.Context, newProduct models.NewProduct, idToken int) error
	GetProductById(ctx context.Context, id int) (models.Product, error)
	GetAllProducts(ctx context.Context) ([]models.Product, error)
	DeleteProduct(ctx context.Context, idToken, id int) error
	UpdateProduct(ctx context.Context, updateProduct models.UpdateProduct, idToken, id int) (models.Category, error)
}

type ProductRepository struct {
	mysql *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		mysql: db,
	}
}

func (pr *ProductRepository) NewProduct(ctx context.Context, newProduct models.NewProduct, idToken int) error {
	query := "INSERT INTO products (name, price, category_id, user_id, quantity, status, images, description, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := pr.mysql.ExecContext(ctx, query, newProduct.Name, newProduct.Price, newProduct.CategoryId, idToken, newProduct.Quantity, newProduct.Status, newProduct.Images, newProduct.Description, newProduct.CreatedAt, newProduct.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProductRepository) GetProductById(ctx context.Context, id int) (models.Product, error)
func (pr *ProductRepository) GetAllProducts(ctx context.Context) ([]models.Product, error)
func (pr *ProductRepository) DeleteProduct(ctx context.Context, idToken, id int) error
func (pr *ProductRepository) UpdateProduct(ctx context.Context, updateProduct models.UpdateProduct, idToken, id int) (models.Category, error)
