package models

import "time"

type User struct {
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Gender    string    `json:"gender" db:"gender"`
	Age       int       `json:"age" db:"age"`
	Address   string    `json:"address" db:"address"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserResponse struct {
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Gender    string    `json:"gender" db:"gender"`
	Age       int       `json:"age" db:"age"`
	Address   string    `json:"address" db:"address"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserUpdate struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Gender   string `json:"gender" db:"gender"`
	Age      int    `json:"age" db:"age"`
	Address  string `json:"address" db:"address"`
}

type UserUpdateResponse struct {
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Gender    string    `json:"gender" db:"gender"`
	Age       int       `json:"age" db:"age"`
	Address   string    `json:"address" db:"address"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
