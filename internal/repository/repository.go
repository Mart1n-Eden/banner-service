package repository

import (
	"banner-service/internal/repository/query"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *Repository {
	return &Repository{db: conn}
}

func (r *Repository) GetUserBanner(tagId, featureId uint64) (string, error) {
	var content string
	if err := r.db.Get(&content, query.GetBanner, tagId, featureId); err != nil {
		return "", err
	}

	return content, nil
}

func (r *Repository) PostBanner(featureId uint64, isActive bool, tagIds []uint64, body json.RawMessage) (uint64, error) {
	var bannerId uint64

	if err := r.db.QueryRow(query.PostBanner, body, isActive).Scan(&bannerId); err != nil {
		// TODO: error handling
		return 0, err
	}

	if _, err := r.db.Exec(query.PostIdentifiers, bannerId, featureId, pq.Array(tagIds)); err != nil {
		// TODO: error handling
		return 0, err
	}

	return bannerId, nil
}
