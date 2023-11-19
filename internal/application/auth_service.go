package application

import (
	"errors"

	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type AuthService struct {
	Repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}

func (s *AuthService) CreateUser(user *domain.User) (*domain.User, error) {

	user, err := s.Repo.SaveUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) VerifyUser(user *domain.User) (string, error) {

	result, err := s.Repo.FindByEmail(user.Email)
	if err != nil {
		return "", errors.New("sorry, the provided email is not registered in our system")
	}

	if err := utils.ComparePassword(result.Password, user.Password); err != nil {
		return "", errors.New("whoops! password mismatch")
	}

	data := map[string]interface{}{
		"id":   result.ID,
	}

	jwt, err := utils.GenerateToken(data)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *AuthService) UpdateBalance(user *domain.User) (*domain.User, error) {

	before, err := s.Repo.FindByID(user.ID)
	if err != nil{
		return nil, err
	}

	user.Balance += before.Balance
	
	if err := s.Repo.UpdateBalance(user); err != nil {
		return nil, err
	}
	
	after, err := s.Repo.FindByID(user.ID)
	if err != nil{
		return nil, err
	}

	return after, nil
}
