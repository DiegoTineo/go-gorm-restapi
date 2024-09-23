package models

import "gorm.io/gorm"

type User struct {
	gorm.Model // esto hara que se lea el struct y se pueda crear la tabla en la BD

	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null;uniqueIndex" json:"email"`
	Tasks     []Task `json:"tasks"`
}
