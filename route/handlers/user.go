package handlers

import (
	"github.com/Lzrb0x/go-gorm-urlShortener-api/route/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	CreateUser(c *gin.Context)
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
