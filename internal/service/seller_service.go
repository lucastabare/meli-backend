package service

import (
	"meli/internal/domain"
	"meli/internal/repository/jsonstore"
)

type SellerService struct{ repo *jsonstore.SellerRepo }
func NewSellerService(r *jsonstore.SellerRepo) *SellerService { return &SellerService{r} }

func (s *SellerService) Get(id string) (domain.Seller, error) {
	return s.repo.GetByID(id)
}
