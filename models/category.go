package models

import "time"

type Category struct {
	Id        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type NewCategory struct {
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UserId    int       `json:"user_id" db:"user_id"`
}

type UpdateCategory struct {
	Name      string     `json:"name" db:"name"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
