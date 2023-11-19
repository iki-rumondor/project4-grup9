package domain

import "time"

type TransactionHistory struct {
	ID          uint `gorm:"primaryKey"`
	Products_Id uint
	User_Id     uint
	Quantity    int
	Total_Price int
	Created_At  time.Time
	Updated_At  time.Time
	Products    Products `gorm:"foreignKey:Products_Id"`
	User        User     `gorm:"foreignKey:User_Id"`
}
