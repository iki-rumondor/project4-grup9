package domain

import "time"

type Categories struct {
	ID                   uint   `gorm:"primaryKey"`
	Type                 string `gorm:"not_nul;varchar(120)"`
	Sold_Product_Ammount int
	Created_At           time.Time
	Updated_At           time.Time
	Product              []Products
}
