package response

import "time"

type DataUsers struct {
	Success bool `json:"success"`
	// Data    []*User `json:"data"`
}

type CreatedUser struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Balance   uint      `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

// type Users struct {
// 	ID        uint   `json:"id"`
// 	Email     string `json:"email"`
// 	Full_Name string `json:"full_name"`
// 	Balance   int    `json:"balance"`

// 	Created_At time.Time `json:"created_at"`
// 	Update_At  time.Time `json:"updated_at"`

// 	Transactions []MyTransaction `json:"transactions"`
// }
