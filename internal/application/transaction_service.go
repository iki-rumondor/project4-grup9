package application

import (
	"errors"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
)

type TransactionService struct {
	Repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		Repo: repo,
	}
}

func (s *TransactionService) CreateTransaction(transaction *domain.TransactionHistory) (*response.TransactionBill, error) {

	product, err := s.Repo.FindProduct(transaction.ProductsId)
	if err != nil {
		return nil, err
	}

	if product.Stock < int(transaction.Quantity) {
		return nil, errors.New("insufficient stock for this product")
	}

	transaction.Total_Price = transaction.Quantity * uint(product.Price)

	user, err := s.Repo.FindUser(transaction.UserId)
	if err != nil {
		return nil, err
	}

	if user.Balance < transaction.Total_Price {
		return nil, errors.New("insufficient balance, please topup your account")
	}

	newBalance := user.Balance - transaction.Total_Price
	newStock := uint(product.Stock) - transaction.Quantity

	if err := s.Repo.CreateTransaction(newBalance, newStock, transaction); err != nil {
		return nil, err
	}

	res := response.TransactionBill{
		Total_Price: transaction.Total_Price,
		Quantity: transaction.Quantity,
		Products_Title: product.Title,
	}

	return &res, nil
}

func (s *TransactionService) GetMyTransaction() (*[]domain.TransactionHistory, error) {
	tasks, err := s.Repo.FindMyTransaction()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TransactionService) GetUserTransaction() (*[]domain.TransactionHistory, error) {
	tasks, err := s.Repo.FindUserTransaction()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
