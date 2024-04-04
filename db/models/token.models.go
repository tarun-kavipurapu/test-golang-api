package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/tarun-kavipurapu/gin-gorm/db"
)

type Token struct {
	TokenID         uint      `gorm:"column:token_id;primaryKey;autoIncrement;not null"`
	TokenValue      string    `gorm:"not null"`
	TokenUserID     uint      `gorm:"column:token_user_id;not null;foreignKey:UserID"`
	TokenExpiryDate time.Time `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func CreateToken(user User) (string, error) {
	var t Token
	t.TokenUserID = user.UserID
	t.TokenValue = uuid.NewString()
	t.TokenExpiryDate = time.Now().Add(2 * 24 * time.Hour)
	err := t.Save()
	if err != nil {
		return "", err
	}
	return t.TokenValue, nil
}
func (t *Token) checkExpired() bool {
	if t.TokenExpiryDate.Before(time.Now()) {
		db.DB.Delete(t.TokenID) //delete the expired token
		return true
	} else {
		return false
	}

}
func (t *Token) Save() error {
	err := db.DB.Save(&t).Error

	if err != nil {
		return err
	}
	return nil
}
