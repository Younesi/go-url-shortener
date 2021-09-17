package entities

import (
	"context"
	"time"
)

type Url struct {
	ID        int64     `json:"id"`
	ShortUrl  string    `json:"short_url"`
	LongUrl   string    `json:"long_url" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// UrlUsecase represent the url's usecases
type UrlUsecase interface {
	GetByShortUrl(ctx context.Context, url string) (string, error)
	Store(ctx context.Context, url string) (string, error)
}
