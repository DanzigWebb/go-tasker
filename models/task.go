package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	UserID     uint
	CategoryID uint

	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	Category    Category  `json:"category"`
}

type TaskApi struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    Category `json:"category"`
	Status      string   `json:"status"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}
