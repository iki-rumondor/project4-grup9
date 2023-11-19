package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type CategoriesRepoImplementation struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) CategoriesRepository {
	return &CategoriesRepoImplementation{
		db: db,
	}
}

func (r *CategoriesRepoImplementation) CreateCategories(categories *domain.Categories) (*domain.Categories, error) {
	if err := r.db.Save(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoriesRepoImplementation) FindCategories() (*[]domain.Categories, error) {
	var category []domain.Categories
	if err := r.db.Preload("Products").Find(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoriesRepoImplementation) UpdateCategories(categories *domain.Categories) (*domain.Categories, error) {
	var result domain.Categories
	if err := r.db.Model(categories).Updates(&categories).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *CategoriesRepoImplementation) DeleteCategories(categories *domain.Categories) error {
	if err := r.db.Delete(&categories).Error; err != nil {
		return err
	}
	return nil
}
