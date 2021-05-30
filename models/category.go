package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	OwnerId     uint
	Title       string `json:"title"`
	Description string `json:"description"`
	Tasks       []Task `gorm:"foreignKey:CategoryID"`
}
