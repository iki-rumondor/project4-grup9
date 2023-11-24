package domain

import (
	"errors"
	"time"

	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primaryKey"`
	FullName    string `gorm:"unique;not_null;varchar(120)"`
	Email       string `gorm:"unique;not_null;varchar(120)"`
	Password    string `gorm:"not_null;varchar(120)"`
	Role      	string `gorm:"not_null;varchar(10)"`
	Balance     uint

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeSave(tx *gorm.DB) error {

	var user User
	if result := tx.First(&user, "email = ? AND id != ?", u.Email, u.ID).RowsAffected; result > 0 {
		return errors.New("the email has already been taken")
	}

	hashPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashPass
	return nil
}

