package application

import (
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

func (s *TransactionService) CreateTransaction(transaction *domain.TransactionHistory) (*domain.TransactionHistory, error) {
	result, err := s.Repo.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	return result, nil
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
