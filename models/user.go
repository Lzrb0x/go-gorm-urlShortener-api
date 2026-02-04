package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	PasswordHash string `json:"-" gorm:"not null"`
	Urls []Url `json:"urls" gorm:"foreignKey:UserID"`
}