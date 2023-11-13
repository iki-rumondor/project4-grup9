package response

import "time"

type Products struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Price         int       `json:"price"`
	Stock         int       `json:"stock"`
	Categories_Id uint      `json:"categories_id"`
	Created_At    time.Time `json:"created_at"`
}

type UpdateProducts struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Price         int       `json:"price"`
	Stock         int       `json:"stock"`
	Categories_Id uint      `json:"categories_id"`
	Created_At    time.Time `json:"created_at"`
	Updated_At    time.Time `json:"updated_at"`
}

type Product struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Price         int       `json:"price"`
	Stock         int       `json:"stock"`
	Categories_Id uint      `json:"categories_id"`
	Created_At    time.Time `json:"created_at"`
	Updated_At    time.Time `json:"updated_at"`

	Transactions []MyTransaction `json:"transactions"`
}
