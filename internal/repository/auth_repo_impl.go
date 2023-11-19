package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type AuthRepoImplementation struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepoImplementation{
		db: db,
	}
}

func (r *AuthRepoImplementation) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepoImplementation) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepoImplementation) SaveUser(user *domain.User) (*domain.User, error) {
	var result domain.User
	if err := r.db.Save(user).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *AuthRepoImplementation) UpdateBalance(user *domain.User) error {
	if err := r.db.Model(&domain.User{}).Where("id = ?", user.ID).Update("balance", user.Balance).Error; err != nil {
		return err
	}

	return nil
}
