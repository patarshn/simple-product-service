package models

import "time"

type Product struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Price        int64     `json:"price"`
	Description  string    `json:"description"`
	CategoryID   int64     `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Stock        int64     `json:"stock"`
	IsAvailable  bool      `json:"is_available"`
	CreatedBy    string    `json:"username"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetAllProductRequest struct {
	Limit  int `json:"limit"`
	Page   int `json:"page"`
	Offset int
}
