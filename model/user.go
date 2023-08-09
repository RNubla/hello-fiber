package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"uniqueIndex;not null" json:"username"`
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	FirstName string `gorm:"not null" json:"firstName"`
	lastName  string `gorm:"not null" json:"lastName"`
}
