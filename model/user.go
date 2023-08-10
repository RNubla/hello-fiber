package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"uniqueIndex;not null" json:"username"`
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	FirstName string `gorm:"not null" json:"firstname"`
	LastName  string `gorm:"not null" json:"lastname"`
}
