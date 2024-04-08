package repository

import (
	"banner-service/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *Repository {
	return &Repository{db: conn}
}

func (r *Repository) GetUserBanner(tag_id, feature_id uint64) (string, error) {
	var content string
	if err := r.db.Get(&content, query.GetUser, tag_id, feature_id); err != nil {
		return "", err
	}

	return content, nil
}
