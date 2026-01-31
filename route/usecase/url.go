package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/Lzrb0x/go-gorm-urlShortener-api/db"
	"github.com/Lzrb0x/go-gorm-urlShortener-api/models"
)
//example of interface usage.
type UrlUsecaseInterface interface {
	GenerateShortURL(originalURL string) (*models.Url, error)
	GetByShortCode(shortCode string) (*models.Url, error)
}

type URLUseCase struct {
	repository db.UrlRepoInterface
}

func NewURLUseCase(repository db.UrlRepoInterface) UrlUsecaseInterface {
	return &URLUseCase{repository: repository}
}

func (u *URLUseCase) GenerateShortURL(originalURL string) (*models.Url, error) {
	existingURL, err := u.repository.GetByOriginalURL(originalURL)
	if err != nil {
		return nil, err
	}
	if existingURL != nil {
		return existingURL, nil
	}

	shortCode, err := u.generateUniqueShortCode()
	if err != nil {
		return nil, err
	}

	newURL := &models.Url{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
	}

	if err := u.repository.Create(newURL); err != nil {
		return nil, err
	}

	return newURL, nil
}

func (u *URLUseCase) GetByShortCode(shortCode string) (*models.Url, error) {
	url, err := u.repository.GetByShortCode(shortCode)
	if err != nil {
		return nil, err
	}

	if url != nil {
		u.repository.IncrementVisits(url)
	}

	return url, nil
}

func (u *URLUseCase) generateUniqueShortCode() (string, error) {
	for i := 0; i < 10; i++ {
		shortCode := u.generateShortCode()

		exists, err := u.repository.CheckShortCodeExists(shortCode)
		if err != nil {
			return "", err
		}

		if !exists {
			return shortCode, nil
		}
	}
	return "", errors.New("failed to generate unique short code")
}

func (u *URLUseCase) generateShortCode() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:8]
}
