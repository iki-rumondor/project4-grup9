package domain

import "time"

type Products struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not_null;varchar(120)"`
	Price       int
	Stock       int
	CategoriesID  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Categories  Categories
	Transaction []TransactionHistory
}
