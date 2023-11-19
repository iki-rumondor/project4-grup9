package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type AuthRepository interface {
	FindByEmail(string) (*domain.User, error)
	FindByID(uint) (*domain.User, error)
	SaveUser(*domain.User) (*domain.User, error)
	UpdateBalance(*domain.User) error
}
