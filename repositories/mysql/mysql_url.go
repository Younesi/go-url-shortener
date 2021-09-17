package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/younesi/go-url-shortener/entities"
	"github.com/younesi/go-url-shortener/repositories"
)

type mysqlUrlRepository struct {
	Conn *sql.DB
}

// NewMysqlUrlRepository will create an implementation of UrlRepository
func NewMysqlUrlRepository(db *sql.DB) repositories.UrlRepository {
	return &mysqlUrlRepository{
		Conn: db,
	}
}

func (m *mysqlUrlRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []entities.Url, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]entities.Url, 0)
	for rows.Next() {
		t := entities.Url{}
		err = rows.Scan(
			&t.LongUrl,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlUrlRepository) GetByShortUrl(ctx context.Context, shortUrl string) (res string, err error) {

	query := `SELECT long_url
  						FROM urls WHERE short_url = ?`

	list, err := m.fetch(ctx, query, shortUrl)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0].LongUrl
	} else {
		return res, errors.New("your requested Item is not found")
	}
	return
}

func (m *mysqlUrlRepository) Store(ctx context.Context, url *entities.Url) (err error) {
	query := `INSERT  urls SET short_url=?, long_url=?, created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, url.ShortUrl, url.LongUrl, url.CreatedAt)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	url.ID = lastID
	return
}
