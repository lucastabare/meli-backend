package jsonstore

import (
	"slices"
	"meli/internal/domain"
)

type ProductRepo struct{ st *Store }
func NewProductRepo(st *Store) *ProductRepo { return &ProductRepo{st: st} }

func (r *ProductRepo) All() ([]domain.Product, error) {
	var items []domain.Product
	if err := r.st.read("products.json", &items); err != nil { return nil, err }
	return items, nil
}

func (r *ProductRepo) GetByID(id string) (domain.Product, error) {
	items, err := r.All()
	if err != nil { return domain.Product{}, err }
	i := slices.IndexFunc(items, func(it domain.Product) bool { return it.ID == id })
	if i == -1 { return domain.Product{}, ErrNotFound }
	return items[i], nil
}

func (r *ProductRepo) SaveAll(items []domain.Product) error {
	return r.st.write("products.json", items)
}
