package usecase

import (
	"time"

	"github.com/Lzrb0x/go-gorm-urlShortener-api/config"
	"github.com/Lzrb0x/go-gorm-urlShortener-api/db"
	"github.com/Lzrb0x/go-gorm-urlShortener-api/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseInterface interface {
	CreateUser(CreateUserRequest *CreateUserRequest) error
	Login(LoginRequest *LoginRequest) (string, string, error)
}

type UserUsecase struct {
	userRepo db.UserRepoInterface
}

func NewUserUsecase(userRepo db.UserRepoInterface) UserUsecaseInterface {
	return &UserUsecase{userRepo: userRepo}
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *UserUsecase) CreateUser(req *CreateUserRequest) error {

	user := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: req.Password,
	}

	return u.userRepo.Create(user)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *UserUsecase) Login(req *LoginRequest) (string, string, error) {

	user, err := u.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", "", err
	}

	accessToken, err := createAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := createRefreshToken(user)
	if err != nil {
		return "", "", err
	}

	user.RefreshToken = refreshToken
	err = u.userRepo.Update(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func createAccessToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Config.JwtSecret))
}

func createRefreshToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Config.JwtSecret))
}
