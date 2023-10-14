package repository

import (
	"database/sql"
	"errors"
	"github.com/riqueGo/url_shortener/contants"
	"github.com/riqueGo/url_shortener/domain"
)

type UrlRepository struct {
	Db *sql.DB
}

func NewUrlRepository(db *sql.DB) *UrlRepository {
	return &UrlRepository{Db: db}
}

func (repo UrlRepository) SaveUrl(urlDomain *domain.UrlDomain) error {

	const maxAttempts = 100

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		urlDomain.GenerateUniqueCode()

		var exists bool
		query := "SELECT EXISTS (SELECT 1 FROM short_url WHERE code = $1)"
		err := repo.Db.QueryRow(query, urlDomain.Code).Scan(&exists)

		if err != nil {
			return err
		}

		if !exists {
			break
		}

		if attempt == maxAttempts {
			return contants.ErrMaxAttemptsReached
		}
	}

	insertQuery := "INSERT INTO short_url (code, url) VALUES ($1, $2)"
	if _, err := repo.Db.Exec(insertQuery, urlDomain.Code, urlDomain.URL); err != nil {
		return err
	}
	return nil
}

func (repo UrlRepository) GetUrl(code string) (string, error) {
	var url string
	query := "SELECT url FROM short_url WHERE code = $1"
	err := repo.Db.QueryRow(query, code).Scan(&url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", contants.ErrCodeNotFound
		}
		return "", err
	}
	return url, nil
}

func (repo UrlRepository) GetAllUrls() ([]domain.UrlDomain, error) {
	rows, err := repo.Db.Query("SELECT code, url FROM short_url")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []domain.UrlDomain
	for rows.Next() {
		var urlDomain domain.UrlDomain
		if err := rows.Scan(&urlDomain.Code, &urlDomain.URL); err != nil {
			return nil, err
		}
		urls = append(urls, urlDomain)
	}
	return urls, nil
}
