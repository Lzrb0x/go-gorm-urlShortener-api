package route

import (
	"github.com/Lzrb0x/go-gorm-urlShortener-api/db"
	"github.com/Lzrb0x/go-gorm-urlShortener-api/route/handlers"
	"github.com/Lzrb0x/go-gorm-urlShortener-api/route/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(database *gorm.DB) *gin.Engine {
	handler := gin.Default()

	urlRepository := db.NewURLRepository(database)
	urlUseCase := usecase.NewURLUseCase(urlRepository)
	urlHandler := handlers.NewURLHandler(urlUseCase)

	handler.POST("/shorten", urlHandler.GenerateShortUrl)
	handler.GET("/:shortCode", urlHandler.RedirectToOriginal)

	return handler
}
