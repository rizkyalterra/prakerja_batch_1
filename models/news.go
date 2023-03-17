package models

import (
	"time"

	"gorm.io/gorm"
)

// struct database
type News struct {
	Id        int            `json:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"deletedAt"`
}

type NewsRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewsResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
