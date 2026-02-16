package handlers

import (
	"github.com/Lzrb0x/go-gorm-urlShortener-api/route/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
}

type UserHandler struct {
	useCase usecase.UserUsecaseInterface
}

func NewUserHandler(useCase usecase.UserUsecaseInterface) UserHandlerInterface {
	return &UserHandler{useCase: useCase}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	req := usecase.CreateUserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.useCase.CreateUser(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully"})
}

func (h *UserHandler) Login(c *gin.Context) {
	req := usecase.LoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.useCase.Login(&req)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
