package config

import (
	"prakerja/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:uV2WvFxfjuI4YJYir73o@tcp(containers-us-west-37.railway.app:6053)/railway?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("tidak koneksi ke database")
	}
	migration()
}

func migration() {
	// automate
	DB.AutoMigrate(&models.News{})
}
