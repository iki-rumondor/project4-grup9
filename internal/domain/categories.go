package domain

import (
	"time"

	"gorm.io/gorm"
)

type Categories struct {
	ID                   uint   `gorm:"primaryKey"`
	Type                 string `gorm:"not_nul;varchar(120)"`
	Sold_Product_Ammount int
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Products             []Products
}

func (m *Categories) BeforeUpdate(tx *gorm.DB) error {

	if err := tx.First(Categories{}, "id = ?", m.ID).Error; err != nil{
		return err
	}

	return nil
}

func (m *Categories) BeforeDelete(tx *gorm.DB) error {

	if err := tx.First(Categories{}, "id = ?", m.ID).Error; err != nil{
		return err
	}

	return nil
}


