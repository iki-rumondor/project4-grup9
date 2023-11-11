package response

import "time"

type Products struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"Title"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Category_Id uint      `json:"category_id"`
	Created_At  time.Time `json:"created_at"`
	Updated_At  time.Time `json:"updated_at"`
}
