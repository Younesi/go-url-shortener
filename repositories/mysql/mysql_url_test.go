package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/younesi/go-url-shortener/entities"
	urlMysqlRepo "github.com/younesi/go-url-shortener/repositories/mysql"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetByShortUrl(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "short_url", "long_url", "created_at"}).
		AddRow(1, "short url 1", "long url 1", time.Now())

	query := "SELECT long_url,created_at FROM urls WHERE short_url = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := urlMysqlRepo.NewMysqlUrlRepository(db)

	shortUrl := "short url 1"

	aUrl, err := a.GetByShortUrl(context.TODO(), shortUrl)
	assert.NoError(t, err)
	assert.NotNil(t, aUrl)
}

func TestStore(t *testing.T) {
	now := time.Now()
	url := &entities.Url{
		ShortUrl:  "Judul",
		LongUrl:   "Content",
		CreatedAt: now,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT urls SET short_url=\\?, long_url=\\?, created_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(url.ShortUrl, url.LongUrl, url.CreatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	a := urlMysqlRepo.NewMysqlUrlRepository(db)

	err = a.Store(context.TODO(), url)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), url.ID)
}
