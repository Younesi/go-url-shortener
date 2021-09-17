package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/younesi/go-url-shortener/controllers"
)

func LoadApiRoutes(route *gin.Engine, handler controllers.UrlHandler) {

	v1 := route.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to the URL Shortener API Test Project",
			})
		})

		v1.POST("/create-short-url", func(c *gin.Context) {
			handler.CreateShortUrl(c)
		})

		v1.GET("/:shortUrl", func(c *gin.Context) {
			handler.HandleShortUrlRedirect(c)
		})
	}

}
