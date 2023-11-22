package response

import "time"

type TransactionBill struct {
	Total_Price    uint    `json:"total_price"`
	Quantity       uint    `json:"quantity"`
	Products_Title string `json:"products_title"`
}

type MyTransaction struct {
	ID          uint `json:"id"`
	ProductsId  uint `json:"products_id"`
	UserId      uint `json:"user_id"`
	Quantity    uint  `json:"quantity"`
	Total_Price uint  `json:"total_price"`
	Product     TransactionProduct
}

type TransactionProduct struct {
	ID           uint            `json:"id"`
	Title        string          `json:"title"`
	Price        int             `json:"price"`
	Stock        int             `json:"stock"`
	CategoriesID uint            `json:"categories_id"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

type TransactionUser struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Balance  uint   `json:"balance"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserTransaction struct {
	ID          uint `json:"id"`
	ProductsId  uint `json:"products_id"`
	UserId      uint `json:"user_id"`
	Quantity    uint  `json:"quantity"`
	Total_Price uint  `json:"total_price"`

	Product TransactionProduct
	Users   TransactionUser
}
