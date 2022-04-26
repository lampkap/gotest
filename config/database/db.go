package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	Amount int    `json:"amount" gorm:"not null"`
}

var DB *gorm.DB

func Open() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/transactions?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))

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
