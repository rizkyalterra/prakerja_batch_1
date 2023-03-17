package config

import (
	"prakerja/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("tidak koneksi ke database")
	}
	migration()
}

func migration() {
	// automate
	DB.AutoMigrate(&models.News{})
}
