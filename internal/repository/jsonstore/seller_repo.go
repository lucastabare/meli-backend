package jsonstore

import "meli/internal/domain"

type SellerRepo struct{ st *Store }
func NewSellerRepo(st *Store) *SellerRepo { return &SellerRepo{st: st} }

func (r *SellerRepo) GetByID(id string) (domain.Seller, error) {
	var sellers []domain.Seller
	if err := r.st.read("sellers.json", &sellers); err != nil { return domain.Seller{}, err }
	for _, s := range sellers {
		if s.ID == id { return s, nil }
	}
	return domain.Seller{}, ErrNotFound
}
