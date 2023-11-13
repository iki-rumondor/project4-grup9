package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type TransactionRepository interface {
	CreateTransaction(*domain.TransactionHistory) (*domain.TransactionHistory, error)
	FindMyTransaction() (*[]domain.TransactionHistory, error)
	FindUserTransaction() (*[]domain.TransactionHistory, error)
}
