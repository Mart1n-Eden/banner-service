package service

import (
	"banner-service/internal/handler/model/request"
	"banner-service/internal/handler/model/response"
	"fmt"
)

type Service interface {
	GetUserBanner(tagId, featureId uint64, useLast bool) (content string, err error)
	GetAdminBanner(tag_id, featureId, limmit, offset *uint64) (res []response.Banner, err error)
	PostBanner(ban request.Banner) (bannerId uint64, err error)
	DeleteBanner(bannerId uint64) error
	PatchBanner(id uint64, ban request.Banner) error
}

type Repository interface {
	GetUserBanner(tagId, featureId uint64) (string, error)
	GetBanner(tag_id, featureId, limmit, offset *uint64) ([]response.Banner, error)
	PostBanner(ban request.Banner) (uint64, error)
	DeleteBanner(bannerId uint64) error
	PatchBanner(id uint64, ban request.Banner) error
}

type Cache interface {
	Set(key string, content string) error
	Get(key string) (content string, err error)
	Exist(key string) bool
}

type service struct {
	repository Repository
	cache      Cache
}

func NewService(repo Repository, c Cache) Service {
	return &service{repository: repo, cache: c}
}

func (s *service) GetUserBanner(tagId, featureId uint64, useLast bool) (content string, err error) {
	key := fmt.Sprintf("%d_%d", tagId, featureId)

	if !useLast {
		if content, err = s.cache.Get(key); err != nil {
			if err.Error() == "not exist" {
			} else {
				return "", err
			}
		} else {
			return content, nil
		}
	}

	if content, err = s.repository.GetUserBanner(tagId, featureId); err != nil {
		return "", err
	}

	if !s.cache.Exist(key) {
		// TODO: error handling
		s.cache.Set(key, content)
	}

	return content, nil
}

func (s *service) GetAdminBanner(tag_id, featureId, limmit, offset *uint64) (res []response.Banner, err error) {
	if res, err = s.repository.GetBanner(tag_id, featureId, limmit, offset); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *service) PostBanner(ban request.Banner) (bannerId uint64, err error) {
	if bannerId, err = s.repository.PostBanner(ban); err != nil {
		return 0, err
	}

	return bannerId, nil
}

func (s *service) DeleteBanner(bannerId uint64) error {
	if err := s.repository.DeleteBanner(bannerId); err != nil {
		return err
	}

	return nil
}

func (s *service) PatchBanner(id uint64, ban request.Banner) error {
	if err := s.repository.PatchBanner(id, ban); err != nil {
		return err
	}

	return nil
}
