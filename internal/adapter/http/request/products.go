package request

type Products struct {
	Title         string `json:"title"`
	Price         int    `json:"price"`
	Stock         int    `json:"stock"`
	Categories_Id uint   `json:"categories_id"`
}
