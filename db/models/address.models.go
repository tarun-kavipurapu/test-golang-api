package models

import (
	"time"

	"github.com/tarun-kavipurapu/gin-gorm/db"
)

type Address struct {
	AddressID uint   `gorm:"column:address_id;primaryKey;not null"`
	Street    string `gorm:"column:address_street"`
	City      string `gorm:"column:address_city"`
	State     string `gorm:"column:address_state"`
	Country   string `gorm:"column:address_country"`
	Zipcode   string `gorm:"column:address_zipcode"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Address) Save() (*Address, error) {
	err := db.DB.Save(&a).Error

	if err != nil {
		return nil, err
	}
	return a, nil
}
