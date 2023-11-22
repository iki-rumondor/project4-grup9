package request

type Transaction struct {
	ProductsId uint `json:"product_id"`
	Quantity   uint  `json:"quantity"`
}
