package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/younesi/go-url-shortener/entities"
)

type UrlRequest struct {
	ID        int64     `json:"id"`
	ShortUrl  string    `json:"short_url"`
	LongUrl   string    `json:"long_url" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type UrlHandler struct {
	urlService entities.UrlUsecase
}

func NewUrlController(service entities.UrlUsecase) (handler UrlHandler) {
	return UrlHandler{
		urlService: service,
	}
}
func (u *UrlHandler) CreateShortUrl(c *gin.Context) {
	var creationRequest UrlRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorg": err.Error()})
		return
	}

	shortUrl, _ := u.urlService.Store(c, creationRequest.LongUrl)

	apiPort := viper.GetString(`database.port`)
	serverPort := viper.GetString(`server.port`)
	host := fmt.Sprintf("http://localhost:%s/api/%s/", apiPort, serverPort)

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

func (u *UrlHandler) HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	initialUrl, err := u.urlService.GetByShortUrl(c, shortUrl)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "Not found!",
		})
	}

	c.Redirect(302, initialUrl)
}
