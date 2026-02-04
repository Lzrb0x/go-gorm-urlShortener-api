package usecase

import (
	"errors"

	"github.com/Lzrb0x/go-gorm-urlShortener-api/db"
	"github.com/Lzrb0x/go-gorm-urlShortener-api/models"
)

type UserUsecaseInterface interface {
	CreateUser(CreateUserRequest *CreateUserRequest) error
}

type UserUsecase struct {
	userRepo db.UserRepoInterface
}

func NewUserUsecase(userRepo db.UserRepoInterface) UserUsecaseInterface {
	return &UserUsecase{userRepo: userRepo}
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *UserUsecase) CreateUser(req *CreateUserRequest) error {
	// Validação de entrada
	if req == nil {
		return errors.New("request não pode ser nil")
	}

	if req.Username == "" {
		return errors.New("username é obrigatório")
	}

	if len(req.Password) < 6 {
		return errors.New("senha deve ter no mínimo 6 caracteres")
	}

	// Criar novo usuário
	user := &models.User{
		Username:     req.Username,
		PasswordHash: req.Password, // Será hashada no repositório
	}

	// Salvar no banco de dados
	return u.userRepo.Create(user)
}
