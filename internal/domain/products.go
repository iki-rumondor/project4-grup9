package domain

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Products struct {
	ID           uint   `gorm:"primaryKey"`
	Title        string `gorm:"not_null;varchar(120)"`
	Price        int
	Stock        int
	CategoriesID uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Categories   Categories
	Transaction  []TransactionHistory
}

func (m *Products) BeforeCreate(tx *gorm.DB) error {

	if err := tx.First(&Categories{}, "id = ?", m.CategoriesID).Error; err != nil {
		return fmt.Errorf("categories with id %d is not found", m.CategoriesID)
	}

	return nil
}

func (m *Products) BeforeDelete(tx *gorm.DB) error {

	if err := tx.First(&Products{}, "id = ?", m.ID).Error; err != nil {
		return err
	}

	return nil
}
