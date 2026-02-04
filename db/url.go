package db

import (
	"errors"

	"github.com/Lzrb0x/go-gorm-urlShortener-api/models"
	"gorm.io/gorm"
)

type UrlRepoInterface interface {
	GetByOriginalURL(originalURL string) (*models.Url, error)
	GetByShortCode(shortCode string) (*models.Url, error)
	Create(url *models.Url) error
	IncrementVisits(url *models.Url) error
	CheckShortCodeExists(shortCode string) (bool, error)
}

type URLRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) UrlRepoInterface {
	return &URLRepository{db: db}
}

func (r *URLRepository) GetByOriginalURL(originalURL string) (*models.Url, error) {
	var url models.Url
	if err := r.db.Where("original_url = ?", originalURL).First(&url).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &url, nil
}

func (r *URLRepository) GetByShortCode(shortCode string) (*models.Url, error) {
	var url models.Url
	if err := r.db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &url, nil
}

func (r *URLRepository) Create(url *models.Url) error {
	return r.db.Create(url).Error
}

func (r *URLRepository) IncrementVisits(url *models.Url) error {
	return r.db.Model(url).Update("visits", gorm.Expr("visits + ?", 1)).Error
}

func (r *URLRepository) CheckShortCodeExists(shortCode string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Url{}).Where("short_code = ?", shortCode).Count(&count).Error
	return count > 0, err
}

