package jsonstore

import "meli/internal/domain"

type PaymentRepo struct{ st *Store }
func NewPaymentRepo(st *Store) *PaymentRepo { return &PaymentRepo{st: st} }

func (r *PaymentRepo) All() ([]domain.PaymentMethod, error) {
	var methods []domain.PaymentMethod
	if err := r.st.read("payments.json", &methods); err != nil { return nil, err }
	return methods, nil
}
