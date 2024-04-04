// models/setup.go

package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	DB = ConnectDatabase()
	return DB
}

func ConnectDatabase() *gorm.DB {
	dsn := "host=db user=postgres password=postgres dbname=testdb port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database

	return DB
}
