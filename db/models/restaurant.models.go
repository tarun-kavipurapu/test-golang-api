package models

import (
	"time"

	"github.com/tarun-kavipurapu/gin-gorm/db"
)

type Restaurant struct {
	RestaurantID        uint    `gorm:"column:restaurant_id;primaryKey;not null"`
	RestaurantName      string  `gorm:"column:restaurant_name"`
	RestaurantUserID    uint    `gorm:"column:restaurant_user_id;not null;foreignKey:UserID"`
	RestaurantAddressID uint    `gorm:"column:restaurant_address_id;not null;foreignKey:AddressID"`
	RestaurantContact   string  `gorm:"column:restaurant_contact"`
	RestaurantRatings   float64 `gorm:"column:restaurant_ratings"`
	RestaurantMinAmount float64 `gorm:"column:restaurant_minAmount"`
	RestaurantMaxAmount float64 `gorm:"column:restaurant_maxAmount"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (r *Restaurant) Save() (*Restaurant, error) {
	err := db.DB.Save(&r).Error

	if err != nil {
		return nil, err
	}
	return r, nil
}
