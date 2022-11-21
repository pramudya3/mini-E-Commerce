package models

type LoginRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginResponse struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Token    string `json:"token" db:"token"`
}

type LoginDataResponse struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
