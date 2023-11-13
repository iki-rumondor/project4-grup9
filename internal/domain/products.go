package domain

import "time"

type Products struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not_null;varchar(120)"`
	Price       int
	Stock       int
	Category_Id uint
	Created_At  time.Time
	Updated_At  time.Time
	Categories  Categories `gorm:"foreignKey:Category_Id"`
	Transaction []TransactionHistory
}
