package response

type TransactionBill struct {
	Total_Price    int    `json:"total_price"`
	Quantity       int    `json:"quantity"`
	Products_Title string `json:"products_title"`
}

type MyTransaction struct {
	ID          uint `json:"id"`
	ProductsId  uint `json:"products_id"`
	UserId      uint `json:"user_id"`
	Quantity    int  `json:"quantity"`
	Total_Price int  `json:"total_price"`
	Product     Product
}

type UserTransaction struct {
	ID          uint `json:"id"`
	ProductsId  uint `json:"products_id"`
	UserId      uint `json:"user_id"`
	Quantity    int  `json:"quantity"`
	Total_Price int  `json:"total_price"`

	Product Product
	Users   Users
}
