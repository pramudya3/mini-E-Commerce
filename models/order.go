package models

import "time"

type Order struct {
	Id            int        `json:"id" db:"id"`
	CartId        int        `json:"cart_id" db:"cart_id"`
	PaymentMethod string     `json:"payment_method" db:"payment_method"`
	StatusOrder   string     `json:"status_order" db:"status_order"`
	PaidAt        *time.Time `json:"paid_at" db:"paid_at"`
	TotalPrice    string     `json:"total_price" db:"total_price"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
}

type NewOrder struct {
	CartId        int    `json:"cart_id" db:"cart_id"`
	PaymentMethod string `json:"payment_method" db:"payment_method"`
	StatusOrder   string `json:"status_order" db:"status_order"`
	TotalPrice    string `json:"total_price" db:"total_price"`
}

type UpdateOrder struct {
	PaidAt *time.Time `json:"paid_at" db:"paid_at"`
}
