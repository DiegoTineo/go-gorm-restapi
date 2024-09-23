package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model // esto hara que se lea el struct y se pueda crear la tabla en la BD

	Title       string `gorm:"not null;unique_index" json:"title"`
	Description string `json:"description"`
	Done        bool   `gorm:"default:false" json:"done"`
	UserID      uint   `json:"user_id"`
}
