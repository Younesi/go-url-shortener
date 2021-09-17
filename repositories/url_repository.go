package repositories

import (
	"context"

	"github.com/younesi/go-url-shortener/entities"
)

// UrlRepository represent the url's repository contract
type UrlRepository interface {
	GetByShortUrl(ctx context.Context, url string) (string, error)
	Store(context.Context, *entities.Url) error
}
