package models

import (
	"time"
)

type Product struct {
	Id          int        `json:"id" db:"id"`
	Name        string     `json:"product_name" db:"product_name"`
	Price       string     `json:"price" db:"price"`
	CategoryId  string     `json:"category_name" db:"category_name"`
	UserId      int        `json:"user_id" db:"user_id"`
	Quantity    int        `json:"quantity" db:"quantity"`
	Status      string     `json:"status" db:"status"`
	Description string     `json:"description" db:"description"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
	// Images      image.Image `json:"images" db:"images"`
}

type NewProduct struct {
	Name        string    `json:"name" db:"name"`
	Price       string    `json:"price" db:"price"`
	CategoryId  int       `json:"category_id" db:"category_id"`
	UserId      int       `json:"user_id" db:"user_id"`
	Quantity    int       `json:"quantity" db:"quantity"`
	Status      string    `json:"status" db:"status"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	// Images      image.Image `json:"images" db:"images"`
}

type UpdateProduct struct {
	Name        string     `json:"name" db:"name"`
	Price       string     `json:"price" db:"price"`
	CategoryId  int        `json:"category_id" db:"category_id"`
	Quantity    int        `json:"quantity" db:"quantity"`
	Status      string     `json:"status" db:"status"`
	Description string     `json:"description" db:"description"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
	// Images      image.Image `json:"images" db:"images"`
}

type ResponseProduct struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Price       string `json:"price" db:"price"`
	CategoryId  string `json:"category_name" db:"category_name"`
	Quantity    int    `json:"quantity" db:"quantity"`
	Status      string `json:"status" db:"status"`
	Description string `json:"description" db:"description"`
	// Images      image.Image `json:"images" db:"images"`
	// CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	// UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}
