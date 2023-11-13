package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type TransactionRepoImplementation struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepoImplementation{
		db: db,
	}
}

func (r *TransactionRepoImplementation) CreateTransaction(transaction *domain.TransactionHistory) (*domain.TransactionHistory, error) {
	if err := r.db.Save(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepoImplementation) FindMyTransaction() (*[]domain.TransactionHistory, error) {
	var transaction []domain.TransactionHistory
	if err := r.db.Preload("Products").Find(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepoImplementation) FindUserTransaction() (*[]domain.TransactionHistory, error) {
	var transaction []domain.TransactionHistory
	if err := r.db.Preload("User").Preload("Products").Find(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}
