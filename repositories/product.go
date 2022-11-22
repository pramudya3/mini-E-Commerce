package repositories

import (
	"context"
	"database/sql"
	"e-commerce/models"
	"errors"
)

type ProductRepositoryInterface interface {
	NewProduct(ctx context.Context, newProduct models.NewProduct, idToken int) error
	GetProductById(ctx context.Context, id int) (models.Product, error)
	GetAllProducts(ctx context.Context) ([]models.ResponseProduct, error)
	DeleteProduct(ctx context.Context, idToken, id int) error
	UpdateProduct(ctx context.Context, updateProduct models.Product, idToken, id int) (models.Product, error)
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
	query := "INSERT INTO products (name, price, category_id, user_id, quantity, status, description) VALUES(?, ?, ?, ?, ?, ?, ?)"

	_, err := pr.mysql.ExecContext(ctx, query, newProduct.Name, newProduct.Price, newProduct.CategoryId, idToken, newProduct.Quantity, newProduct.Status, newProduct.Description)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProductRepository) GetProductById(ctx context.Context, id int) (models.Product, error) {
	var product models.Product
	query := "SELECT p.id, p.name, p.price, c.name as category_name, p.quantity, p.status, p.description FROM products p LEFT JOIN categories c on c.id = p.category_id WHERE p.id = ?"

	err := pr.mysql.QueryRowContext(ctx, query, id).Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.Quantity, &product.Status, &product.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, errors.New("data not found")
		}
		return models.Product{}, err
	}
	return product, nil
}

func (pr *ProductRepository) GetAllProducts(ctx context.Context) ([]models.ResponseProduct, error) {
	query := "SELECT p.id, p.name, p.price, c.name as category_name, p.quantity, p.status, p.description FROM products p LEFT JOIN categories c on c.id = p.category_id"

	rows, err := pr.mysql.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.ResponseProduct
	for rows.Next() {
		var product models.ResponseProduct
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.Quantity, &product.Status, &product.Description)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (pr *ProductRepository) DeleteProduct(ctx context.Context, idToken, id int) error {
	query := "DELETE FROM products WHERE user_id = ? AND id = ?"

	result, err := pr.mysql.ExecContext(ctx, query, idToken, id)
	if err != nil {
		return nil
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (pr *ProductRepository) UpdateProduct(ctx context.Context, updateProduct models.Product, idToken, id int) (models.Product, error) {
	query := "UPDATE products SET name = ?, price = ?, category_id = ?, quantity = ?, status = ?, description = ?, updated_at = ? WHERE user_id = ? AND id = ?"

	result, err := pr.mysql.ExecContext(ctx, query, updateProduct.Name, updateProduct.Price, updateProduct.CategoryId, updateProduct.Quantity, updateProduct.Status, updateProduct.Description, updateProduct.UpdatedAt, idToken, id)
	if err != nil {
		return models.Product{}, err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return models.Product{}, err
	}
	return updateProduct, nil
}
