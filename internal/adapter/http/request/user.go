package request

type Register struct {
	FullName string `json:"full_name" valid:"required~field full_name is required"`
	Email    string `json:"email" valid:"required~field email is required, email"`
	Password string `json:"password" valid:"required~field password is required, length(6|99)~password at least 6 character"`
}

type Login struct {
	Email    string `json:"email" valid:"required~please make sure to provide password in the request, email"`
	Password string `json:"password" valid:"required~please make sure to provide password in the request "`
}

type Topup struct {
	Balance uint `json:"balance" valid:"range(0|100000000)~top up minimum balance 0 and maximum 100000000"`
}
