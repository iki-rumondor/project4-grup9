package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type CategoriesRepository interface {
	CreateCategories(*domain.Categories) (*domain.Categories, error)
	FindCategories() (*[]domain.Categories, error)
	UpdateCategories(*domain.Categories) (*domain.Categories, error)
	DeleteCategories(*domain.Categories) error
}
