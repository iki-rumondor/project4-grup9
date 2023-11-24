package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type TransactionRepository interface {
	FindProduct(uint) (*domain.Products, error)
	FindUser(uint) (*domain.User, error)
	CreateTransaction(newBalance, newStock uint, transaction *domain.TransactionHistory) error
	FindMyTransaction() (*[]domain.TransactionHistory, error)
	FindUserTransaction() (*[]domain.TransactionHistory, error)
}
