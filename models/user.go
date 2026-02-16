package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique;not null;index"`
	PasswordHash string `json:"-" gorm:"not null"`
	RefreshToken string `json:"refresh_token"`
	Urls         []Url  `json:"urls" gorm:"foreignKey:UserID"`
}
