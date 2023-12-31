package response

import "time"

type Categories struct {
	ID                   uint                  `gorm:"primaryKey" json:"id"`
	Type                 string                `json:"type"`
	Sold_Product_Ammount int                   `json:"sold_product_ammount"`
	CreatedAt            time.Time             `json:"created_at"`
	UpdatedAt            time.Time             `json:"updated_at"`
	ProductsCategories   []*ProductsCategories `json:"products"`
}

type CreateCategories struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	Type                 string    `json:"type"`
	Sold_Product_Ammount int       `json:"sold_product_ammount"`
	Created_At           time.Time `json:"created_at"`
}

type UpdateCategories struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	Type                 string    `json:"type"`
	Sold_Product_Ammount int       `json:"sold_product_ammount"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type ProductsCategories struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `json:"Title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryId uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
