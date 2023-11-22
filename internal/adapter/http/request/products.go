package request

type Products struct {
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoriesID uint   `json:"categories_id"`
}
