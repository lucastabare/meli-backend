package jsonstore

import (
	"errors"
	"io/fs"
	"meli/internal/domain"
)

type AdsRepo struct{ st *Store }

func NewAdsRepo(st *Store) *AdsRepo {
	return &AdsRepo{st}
}

func (r *AdsRepo) List() ([]domain.Ad, error) {
	var out []domain.Ad
	if err := r.st.read("ads.json", &out); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return []domain.Ad{}, nil
		}
		return nil, err
	}
	return out, nil
}
