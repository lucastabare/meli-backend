package service

import "meli/internal/domain"

type AdsRepository interface {
	List() ([]domain.Ad, error)
}

type AdsService struct{ repo AdsRepository }

func NewAdsService(r AdsRepository) *AdsService { return &AdsService{r} }

func (s *AdsService) List() ([]domain.Ad, error) {
	return s.repo.List()
}
