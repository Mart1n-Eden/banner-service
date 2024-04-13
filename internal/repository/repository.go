package repository

import (
	"banner-service/internal/handler/model/request"
	"banner-service/internal/handler/model/response"
	"banner-service/internal/repository/query"
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
	var isActive bool
	if err := r.db.QueryRow(query.GetBanner, tagId, featureId).Scan(&content, &isActive); err != nil {
		return "", err
	}

	if !isActive {
		return "", nil
	}

	return content, nil
}

func (r *Repository) GetBanner(tag_id, featureId, limmit, offset *uint64) ([]response.Banner, error) {
	var res []response.Banner

	if rows, err := r.db.Queryx(query.GetAdminBanner, tag_id, featureId, limmit, offset); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var ban response.Banner
			if err := rows.Scan(&ban.Id, pq.Array(&ban.TagsId), &ban.FeatureId, &ban.Content, &ban.IsActive, &ban.Created, &ban.Updated); err != nil {
				return nil, err
			}
			res = append(res, ban)
		}
	}

	return res, nil
}

func (r *Repository) PostBanner(ban request.Banner) (uint64, error) {
	var bannerId uint64

	if err := r.db.QueryRow(query.PostBanner, *ban.Content, *ban.IsActive).Scan(&bannerId); err != nil {
		// TODO: error handling
		return 0, err
	}

	if _, err := r.db.Exec(query.PostIdentifiers, bannerId, *ban.FeatureId, pq.Array(ban.TagIds)); err != nil {
		// TODO: error handling
		return 0, err
	}

	return bannerId, nil
}

func (r *Repository) DeleteBanner(bannerId uint64) error {

	if _, err := r.db.Exec(query.DeleteBanner, bannerId); err != nil {
		// TODO: error handling
		return err
	}

	return nil

}

// TODO: implemetation with transaction
func (r *Repository) PatchBanner(id uint64, ban request.Banner) error {

	if ban.Content != nil {
		if _, err := r.db.Exec(query.PatchContent, *ban.Content, id); err != nil {
			// TODO: error handling
			return err
		}
	}

	if ban.IsActive != nil {
		if _, err := r.db.Exec(query.PatchIsActive, *ban.IsActive, id); err != nil {
			// TODO: error handling
			return err
		}
	}

	switch {
	case ban.TagIds == nil && ban.FeatureId != nil:
		if _, err := r.db.Exec(query.PatchFeature, *ban.FeatureId, id); err != nil {
			// TODO: error handling
			return err
		}

	case ban.TagIds != nil && ban.FeatureId != nil:
		if _, err := r.db.Exec(query.DeleteIdentifier, id); err != nil {
			// TODO: error handling
			return err
		}

		if _, err := r.db.Exec(query.PostIdentifiers, id, *ban.FeatureId, pq.Array(ban.TagIds)); err != nil {
			// TODO: error handling
			return err
		}

	case ban.TagIds != nil && ban.FeatureId == nil:
		var featureId uint64

		if err := r.db.Get(&featureId, query.SelectFeature, id); err != nil {
			return err
		}

		if _, err := r.db.Exec(query.DeleteIdentifier, id); err != nil {
			// TODO: error handling
			return err
		}

		if _, err := r.db.Exec(query.PostIdentifiers, id, featureId, pq.Array(ban.TagIds)); err != nil {
			// TODO: error handling
			return err
		}

	}

	return nil
}
