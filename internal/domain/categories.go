package domain

import "time"

type Categories struct {
	ID                   uint   `gorm:"primaryKey"`
	Type                 string `gorm:"not_nul;varchar(120)"`
	Sold_Product_Ammount int
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Product              []Products `gorm:"foreignKey:CategoryId"`
}
