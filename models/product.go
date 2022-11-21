package models

import (
	"image"
	"time"
)

type Product struct {
	Id          int         `json:"id" db:"id"`
	Name        string      `json:"product_name" db:"product_name"`
	Price       string      `json:"price" db:"price"`
	CategoryId  int         `json:"category_id" db:"category_id"`
	UserId      int         `json:"user_id" db:"user_id"`
	Quantity    int         `json:"quantity" db:"quantity"`
	Status      string      `json:"status" db:"status"`
	Images      image.Image `json:"images" db:"images"`
	Description string      `json:"description" db:"description"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at" db:"updated_at"`
}

type NewProduct struct {
	Name        string      `json:"product_name" db:"product_name"`
	Price       string      `json:"price" db:"price"`
	CategoryId  int         `json:"category_id" db:"category_id"`
	UserId      int         `json:"user_id" db:"user_id"`
	Quantity    int         `json:"quantity" db:"quantity"`
	Status      string      `json:"status" db:"status"`
	Images      image.Image `json:"images" db:"images"`
	Description string      `json:"description" db:"description"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at" db:"updated_at"`
}

type UpdateProduct struct {
	Name        string      `json:"product_name" db:"product_name"`
	Price       string      `json:"price" db:"price"`
	CategoryId  int         `json:"category_id" db:"category_id"`
	UserId      int         `json:"user_id" db:"user_id"`
	Quantity    int         `json:"quantity" db:"quantity"`
	Status      string      `json:"status" db:"status"`
	Images      image.Image `json:"images" db:"images"`
	Description string      `json:"description" db:"description"`
	UpdatedAt   *time.Time  `json:"updated_at" db:"updated_at"`
}
