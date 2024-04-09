package service

import (
	"banner-service/internal/handler/model/request"
	"banner-service/internal/repository"
	"fmt"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repository: repo}
}

func (s *Service) GetUserBanner(tag_id, feature_id uint64) (content string, err error) {
	if content, err = s.repository.GetUserBanner(tag_id, feature_id); err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return content, nil
}

//func (s *Service) GetAdminBanner()

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
