package repository

import (
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type ProductsRepoImplementation struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) ProductsRepository {
	return &ProductsRepoImplementation{
		db: db,
	}
}

func (r *ProductsRepoImplementation) CreateProducts(products *domain.Products) (*domain.Products, error) {
	if err := r.db.Save(products).First(products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductsRepoImplementation) FindProducts() (*[]domain.Products, error) {
	var products []domain.Products
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return &products, nil
}

func (r *ProductsRepoImplementation) UpdateProducts(products *domain.Products) (*domain.Products, error) {
	if err := r.db.First(&domain.Categories{}, "id = ?", products.CategoriesID).Error; err != nil {
		return nil, fmt.Errorf("categories with id %d is not found", products.CategoriesID)
	}

	var result domain.Products
	if err := r.db.Model(products).Updates(products).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ProductsRepoImplementation) DeleteProducts(products *domain.Products) error {
	if err := r.db.Delete(products).Error; err != nil {
		return err
	}
	return nil
}
