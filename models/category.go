package models

type CategoryRequest struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
