package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	OriginalURL string `json:"original_url" gorm:"not null"`
	ShortCode   string `json:"short_code" gorm:"unique;not null;index"`
	Visits      int64  `json:"visits" gorm:"default:0"`
}
