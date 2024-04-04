package models

import (
	"errors"
	"strings"
	"time"

	"github.com/tarun-kavipurapu/gin-gorm/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User struct
type User struct {
	UserID       uint   `gorm:"column:user_id;primaryKey;not null"`
	UserEmail    string `gorm:"column:user_email;uniqueIndex;not null"`
	UserPassword string `gorm:"column:user_password;not null"`
	UserRoleID   uint   `gorm:"column:user_role_id;not null;foreignKey:RoleID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) Save() (*User, error) {
	err := db.DB.Save(&u).Error

	if err != nil {
		return nil, err
	}
	return u, nil
}

// hash the password before saving into the databse
func (u *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.UserPassword = string(passwordHash)

	return nil

}

//validate password

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.UserPassword), []byte(password))

}

//Get all Users

func GetUserById(id uint) (User, error) {
	var user User
	err := db.DB.Where("user_id=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func GetUserByEmail(email string) (User, error) {
	// Trim leading and trailing whitespaces from the email
	email = strings.TrimSpace(email)

	// Check if the email is empty after trimming
	if email == "" {
		return User{}, errors.New("empty email is not valid")
	}

	var user User
	err := db.DB.Where("user_email = ?", email).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// func (u *User) GetAllUsers() ([]User, error) {

// }
