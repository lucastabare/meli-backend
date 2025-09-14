package service

import (
	"meli/internal/domain"
	"meli/internal/repository/jsonstore"
)

type PaymentService struct{ repo *jsonstore.PaymentRepo }
func NewPaymentService(r *jsonstore.PaymentRepo) *PaymentService { return &PaymentService{r} }
func (s *PaymentService) Methods() ([]domain.PaymentMethod, error) { return s.repo.All() }
