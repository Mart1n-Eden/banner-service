package service

import (
	"banner-service/internal/repository"
	"fmt"
)

type Service struct {
	Repository *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Repository: repo}
}

func (s *Service) GetUserBanner(tag_id, feature_id uint64) (content string, err error) {
	if content, err = s.Repository.GetUserBanner(tag_id, feature_id); err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return content, nil
}

//func (s *Service) GetAdminBanner()

func (s *Service) PostBanner(feature_id uint64, is_active bool, tag_ids []uint64, body string) {

}
