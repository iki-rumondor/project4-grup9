package domain

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TransactionHistory struct {
	ID          uint `gorm:"primaryKey"`
	ProductsId  uint
	UserId      uint
	Quantity    uint
	Total_Price uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Products    Products `gorm:"foreignKey:ProductsId"`
	User        User     `gorm:"foreignKey:UserId"`
}

func (m *TransactionHistory) BeforeSave(tx *gorm.DB) error {

	if err := tx.First(&Products{}, "id = ?", m.ProductsId).Error; err != nil {
		return fmt.Errorf("product with id %d is not found", m.ProductsId)
	}

	if err := tx.First(&User{}, "id = ?", m.UserId).Error; err != nil {
		return fmt.Errorf("user with id %d is not found", m.UserId)
	}

	return nil
}

func (m *TransactionHistory) BeforeDelete(tx *gorm.DB) error {

	if err := tx.First(&TransactionHistory{}, "id = ?", m.ID).Error; err != nil {
		return err
	}

	return nil
}
