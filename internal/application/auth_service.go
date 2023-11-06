package application

import (
	"errors"

	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
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

func (s *AuthService) CreateUser(request *request.Register) error {
	user := &domain.User{
		Uuid:     uuid.NewString(),
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		RoleID:   request.RoleID,
	}

	if err := s.Repo.SaveUser(user); err != nil{
		return err
	}

	return nil
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
		"role": result.RoleID,
	}

	jwt, err := utils.GenerateToken(data)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *AuthService) GetUsers() (*[]domain.User, error) {

	users, err := s.Repo.FindUsers()
	
	if err != nil{
		return nil, err
	}

	return users, nil
}