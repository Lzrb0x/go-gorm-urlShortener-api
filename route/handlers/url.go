package handlers

import (
	"github.com/Lzrb0x/go-gorm-urlShortener-api/route/usecase"
	"github.com/gin-gonic/gin"
)

type generateShortUrlRequest struct {
	OriginalUrl string `json:"original_url" binding:"required,url"`
}

type UrlHandlerInterface interface {
	GenerateShortUrl(c *gin.Context)
	RedirectToOriginal(c *gin.Context)
}

type URLHandler struct {
	useCase usecase.UrlUsecaseInterface // example of receiving interface of usecase
}

func NewURLHandler(useCase usecase.UrlUsecaseInterface) UrlHandlerInterface {
	return &URLHandler{useCase: useCase}
}

func (h *URLHandler) GenerateShortUrl(c *gin.Context) {
	var req generateShortUrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	url, err := h.useCase.GenerateShortURL(req.OriginalUrl)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate short URL"})
		return
	}

	c.JSON(201, gin.H{
		"short_code":   url.ShortCode,
		"original_url": url.OriginalURL,
		"short_url":    c.Request.Host + "/" + url.ShortCode,
	})
}

func (h *URLHandler) RedirectToOriginal(c *gin.Context) {
	shortCode := c.Param("shortCode")

	url, err := h.useCase.GetByShortCode(shortCode)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve URL"})
		return
	}

	if url == nil {
		c.JSON(404, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(302, url.OriginalURL)
}
