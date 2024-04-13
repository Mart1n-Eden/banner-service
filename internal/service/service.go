package service

import (
	"banner-service/internal/cache"
	"banner-service/internal/handler/model/request"
	"banner-service/internal/handler/model/response"
	"banner-service/internal/repository"
	"fmt"
)

type Service struct {
	repository *repository.Repository
	cache      *cache.Cache
}

func NewService(repo *repository.Repository, c *cache.Cache) *Service {
	return &Service{repository: repo, cache: c}
}

func (s *Service) GetUserBanner(tagId, featureId uint64, useLast bool) (content string, err error) {
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

func (s *Service) GetAdminBanner(tag_id, featureId, limmit, offset *uint64) (res []response.Banner, err error) {
	if res, err = s.repository.GetBanner(tag_id, featureId, limmit, offset); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) PostBanner(ban request.Banner) (bannerId uint64, err error) {
	if bannerId, err = s.repository.PostBanner(ban); err != nil {
		return 0, err
	}

	return bannerId, nil
}

func (s *Service) DeleteBanner(bannerId uint64) error {
	if err := s.repository.DeleteBanner(bannerId); err != nil {
		return err
	}

	return nil
}

func (s *Service) PatchBanner(id uint64, ban request.Banner) error {
	if err := s.repository.PatchBanner(id, ban); err != nil {
		return err
	}

	return nil
}
