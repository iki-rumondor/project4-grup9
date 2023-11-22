package response

import "time"

type Products struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Price        int       `json:"price"`
	Stock        int       `json:"stock"`
	CategoriesID uint      `json:"categories_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type UpdateProducts struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Price        string    `json:"price"`
	Stock        int       `json:"stock"`
	CategoriesID uint      `json:"categories_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Product struct {
	ID           uint            `json:"id"`
	Title        string          `json:"title"`
	Price        int             `json:"price"`
	Stock        int             `json:"stock"`
	CategoriesID uint            `json:"categories_id"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	Transactions []MyTransaction `json:"transactions"`
}
