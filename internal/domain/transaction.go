package domain

import "time"

type TransactionHistory struct {
	ID          uint `gorm:"primaryKey"`
	ProductsId  uint
	UserId      uint
	Quantity    int
	Total_Price int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Products    Products `gorm:"foreignKey:ProductsId"`
	User        User     `gorm:"foreignKey:UserId"`
}
