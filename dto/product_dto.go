package dto

import "time"

type ProductCreate struct {
	Name  string  `json:"name" binding:"required" valid:"required~O nome do produto não pode ser vazio"`
	Price float64 `json:"price" binding:"required" valid:"required~O preço é obrigatório"`
}

type ProductUpdate struct {
	Name  *string  `json:"name" valid:"required~O nome do produto não pode ser vazio"`
	Price *float64 `json:"price" valid:"required~O preço é obrigatório"`
}

type ProductResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
