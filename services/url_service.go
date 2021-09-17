package services

import (
	"context"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/younesi/go-url-shortener/entities"
	"github.com/younesi/go-url-shortener/repositories"
	shortener "github.com/younesi/go-url-shortener/shortner"
)

type urlService struct {
	urlRepo        repositories.UrlRepository
	contextTimeout time.Duration
}

// NewUrlservice will create new an urlService object representation of entities.UrlUsecase interface
func NewUrlservice(u repositories.UrlRepository, timeout time.Duration) entities.UrlUsecase {
	return &urlService{
		urlRepo:        u,
		contextTimeout: timeout,
	}
}

func (u *urlService) GetByShortUrl(c context.Context, shortUrl string) (url string, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	url, err = u.urlRepo.GetByShortUrl(ctx, shortUrl)

	if err != nil {
		log.Warn("your requested Item is not found")
	}

	return
}

func (u *urlService) Store(c context.Context, initialUrl string) (shortUrl string, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	shortUrl = shortener.GenerateShortLink(initialUrl)

	now := time.Now()
	url := &entities.Url{
		ShortUrl:  shortUrl,
		LongUrl:   initialUrl,
		CreatedAt: now,
	}
	err = u.urlRepo.Store(ctx, url)
	return shortUrl, err
}
