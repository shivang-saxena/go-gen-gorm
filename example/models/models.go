package models

import "gorm.io/gorm"

//go:generate go-gen-gorm -in models.go

// User is a usual user
// gen:qs
type User struct {
	gorm.Model

	Name    string
	Surname *string `gorm:"column:user_surname"`
	Email   string
}
