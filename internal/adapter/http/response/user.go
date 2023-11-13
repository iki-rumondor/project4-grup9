package response

import "time"

type DataUsers struct {
	Success bool    `json:"success"`
	Data    []*User `json:"data"`
}

type User struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type Users struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Full_Name string `json:"full_name"`
	Balance   int    `json:"balance"`

	Created_At time.Time `json:"created_at"`
	Update_At  time.Time `json:"updated_at"`

	Transactions []MyTransaction `json:"transactions"`
}
