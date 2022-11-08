package models

import "time"

type User struct {
	Id        int        `json:"id" db:"id"`
	Username  string     `json:"username" db:"username"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"password" db:"password"`
	Gender    string     `json:"gender" db:"gender"`
	Age       int        `json:"age" db:"age"`
	Address   string     `json:"address" db:"address"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type CreateUserRequest struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Gender   string `json:"gender" db:"gender"`
	Age      int    `json:"age" db:"age"`
	Address  string `json:"address" db:"address"`
}

type UserResponse struct {
	Id        int        `json:"id" db:"id"`
	Username  string     `json:"username" db:"username"`
	Email     string     `json:"email" db:"email"`
	Gender    string     `json:"gender" db:"gender"`
	Age       int        `json:"age" db:"age"`
	Address   string     `json:"address" db:"address"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type UserUpdateRequest struct {
	Username  string     `json:"username" db:"username"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"password" db:"password"`
	Gender    string     `json:"gender" db:"gender"`
	Age       int        `json:"age" db:"age"`
	Address   string     `json:"address" db:"address"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
