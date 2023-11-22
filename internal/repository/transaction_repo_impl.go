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

func (r *TransactionRepoImplementation) FindMyTransaction() (*[]domain.TransactionHistory, error) {
	var transaction []domain.TransactionHistory
	if err := r.db.Preload("Products").Find(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepoImplementation) CreateTransaction(newBalance, newStock uint, transaction *domain.TransactionHistory) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&domain.Products{}).Where("id = ?", transaction.ProductsId).Update("stock", newStock).Error; err != nil{
			return err
		}
		
		if err := tx.Model(&domain.User{}).Where("id = ?", transaction.UserId).Update("balance", newBalance).Error; err != nil{
			return err
		}

		if err := tx.Save(transaction).Error; err != nil{
			return err
		}
		
		return nil
	})
}

func (r *TransactionRepoImplementation) FindProduct(id uint) (*domain.Products, error) {
	var product domain.Products
	if err := r.db.First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *TransactionRepoImplementation) FindUser(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *TransactionRepoImplementation) FindUserTransaction() (*[]domain.TransactionHistory, error) {
	var transaction []domain.TransactionHistory
	if err := r.db.Preload("User").Preload("Products").Find(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}
