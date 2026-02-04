package db

import (
	"github.com/Lzrb0x/go-gorm-urlShortener-api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	Create(user *models.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepoInterface {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	// Hash da senha com custo 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), 10)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)

	if result := r.db.Create(user); result.Error != nil {
		return result.Error
	}

	return nil
}
