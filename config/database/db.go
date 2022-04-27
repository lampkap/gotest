package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	Amount int    `json:"amount" gorm:"not null"`
}

var DB *gorm.DB

func Open() (err error) {
	dsn := os.Getenv("DATABASE_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB.AutoMigrate(&Transaction{})

	return nil
}

func Close() error {
	sqlDB, err := DB.DB()

	if err != nil {
		return err
	}

	return sqlDB.Close()
}
