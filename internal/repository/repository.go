package repository

import (
	"banner-service/internal/handler/model/request"
	"banner-service/internal/handler/model/response"
	"banner-service/internal/repository/query"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Repository interface {
	GetUserBanner(tagId, featureId uint64) (string, error)
	GetBanner(tag_id, featureId, limmit, offset *uint64) ([]response.Banner, error)
	PostBanner(ban request.Banner) (uint64, error)
	DeleteBanner(bannerId uint64) error
	PatchBanner(id uint64, ban request.Banner) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(conn *sqlx.DB) Repository {
	return &repository{db: conn}
}

func (r *repository) GetUserBanner(tagId, featureId uint64) (string, error) {
	var content string
	var isActive bool
	if err := r.db.QueryRow(query.GetBanner, tagId, featureId).Scan(&content, &isActive); err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	if !isActive {
		return "", nil
	}

	return content, nil
}

func (r *repository) GetBanner(tag_id, featureId, limmit, offset *uint64) ([]response.Banner, error) {
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

func (r *repository) PostBanner(ban request.Banner) (uint64, error) {
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

func (r *repository) DeleteBanner(bannerId uint64) error {

	if _, err := r.db.Exec(query.DeleteBanner, bannerId); err != nil {
		// TODO: error handling
		return err
	}

	return nil

}

// TODO: implemetation with transaction
// func (r *repository) PatchBanner(id uint64, ban request.Banner) error {
//
//		if ban.Content != nil {
//			if _, err := r.db.Exec(query.PatchContent, *ban.Content, id); err != nil {
//				// TODO: error handling
//				return err
//			}
//		}
//
//		if ban.IsActive != nil {
//			if _, err := r.db.Exec(query.PatchIsActive, *ban.IsActive, id); err != nil {
//				// TODO: error handling
//				return err
//			}
//		}
//
//		switch {
//		case ban.TagIds == nil && ban.FeatureId != nil:
//			if _, err := r.db.Exec(query.PatchFeature, *ban.FeatureId, id); err != nil {
//				// TODO: error handling
//				return err
//			}
//
//		case ban.TagIds != nil && ban.FeatureId != nil:
//			if _, err := r.db.Exec(query.DeleteIdentifier, id); err != nil {
//				// TODO: error handling
//				return err
//			}
//
//			if _, err := r.db.Exec(query.PostIdentifiers, id, *ban.FeatureId, pq.Array(ban.TagIds)); err != nil {
//				// TODO: error handling
//				return err
//			}
//
//		case ban.TagIds != nil && ban.FeatureId == nil:
//			var featureId uint64
//
//			if err := r.db.Get(&featureId, query.SelectFeature, id); err != nil {
//				return err
//			}
//
//			if _, err := r.db.Exec(query.DeleteIdentifier, id); err != nil {
//				// TODO: error handling
//				return err
//			}
//
//			if _, err := r.db.Exec(query.PostIdentifiers, id, featureId, pq.Array(ban.TagIds)); err != nil {
//				// TODO: error handling
//				return err
//			}
//
//		}
//
//		return nil
//	}
func (r *repository) PatchBanner(id uint64, ban request.Banner) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			err = fmt.Errorf("%v", p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	if ban.Content != nil {
		if _, err := tx.Exec(query.PatchContent, *ban.Content, id); err != nil {
			return err
		}
	}

	if ban.IsActive != nil {
		if _, err := tx.Exec(query.PatchIsActive, *ban.IsActive, id); err != nil {
			return err
		}
	}

	switch {
	case ban.TagIds == nil && ban.FeatureId != nil:
		if _, err := tx.Exec(query.PatchFeature, *ban.FeatureId, id); err != nil {
			return err
		}

	case ban.TagIds != nil && ban.FeatureId != nil:
		if _, err := tx.Exec(query.DeleteIdentifier, id); err != nil {
			return err
		}

		if _, err := tx.Exec(query.PostIdentifiers, id, *ban.FeatureId, pq.Array(ban.TagIds)); err != nil {
			return err
		}

	case ban.TagIds != nil && ban.FeatureId == nil:
		var featureId uint64
		if err := tx.QueryRow(query.SelectFeature, id).Scan(&featureId); err != nil {
			return err
		}

		if _, err := tx.Exec(query.DeleteIdentifier, id); err != nil {
			return err
		}

		if _, err := tx.Exec(query.PostIdentifiers, id, featureId, pq.Array(ban.TagIds)); err != nil {
			return err
		}
	}

	return nil
}
