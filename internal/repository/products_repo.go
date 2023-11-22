package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type ProductsRepository interface {
	CreateProducts(*domain.Products) (*domain.Products, error)
	FindProducts() (*[]domain.Products, error)
	UpdateProducts(*domain.Products) (*domain.Products, error)
	DeleteProducts(*domain.Products) error
}
