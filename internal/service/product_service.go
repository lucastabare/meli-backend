package service

import (
	"errors"
	"math"
	"sort"
	"strings"
	"meli/internal/domain"
	"meli/internal/repository/jsonstore"
)

type ProductService struct{ repo *jsonstore.ProductRepo }
func NewProductService(r *jsonstore.ProductRepo) *ProductService { return &ProductService{repo: r} }

func (s *ProductService) Get(id string) (domain.Product, error) { return s.repo.GetByID(id) }

type ListParams struct {
	Query  string
	Limit  int
	Offset int
}

func (s *ProductService) List(p ListParams) ([]domain.Product, int, error) {
	items, err := s.repo.All()
	if err != nil {
		if errors.Is(err, jsonstore.ErrNotFound) {
			return []domain.Product{}, 0, nil
		}
		return nil, 0, err
	}

	var filtered []domain.Product
	if p.Query != "" {
		q := strings.ToLower(p.Query)
		for _, it := range items {
			if strings.Contains(strings.ToLower(it.Title), q) ||
				strings.Contains(strings.ToLower(it.Description), q) {
				filtered = append(filtered, it)
			}
		}
	} else {
		filtered = items
	}
	total := len(filtered)

	start := p.Offset
	if start > total { return []domain.Product{}, total, nil }
	end := start + p.Limit
	if p.Limit <= 0 || end > total { end = total }
	return filtered[start:end], total, nil
}

func (s *ProductService) Similar(id string, limit int) ([]domain.Product, error) {
	base, err := s.repo.GetByID(id)
	if err != nil { return nil, err }

	all, err := s.repo.All()
	if err != nil {
		if errors.Is(err, jsonstore.ErrNotFound) { return []domain.Product{}, nil }
		return nil, err
	}

	type scored struct {
		P     domain.Product
		Score float64
	}

	var bucket []scored
	for _, p := range all {
		if p.ID == base.ID { continue }

		score := 0.0
		if p.Category == base.Category { score += 2 }     
		if p.Brand == base.Brand { score += 1.5 }         
		if base.Price > 0 {
			diff := math.Abs(p.Price-base.Price) / base.Price
			if diff <= 0.2 { score += 1 }
		}
		score += float64(sharedCount(base.Tags, p.Tags)) * 0.3

		if score > 0 {
			bucket = append(bucket, scored{P: p, Score: score})
		}
	}

	sort.Slice(bucket, func(i, j int) bool {
		if bucket[i].Score == bucket[j].Score {
			return bucket[i].P.RatingAvg > bucket[j].P.RatingAvg
		}
		return bucket[i].Score > bucket[j].Score
	})

	if limit <= 0 || limit > len(bucket) { limit = len(bucket) }
	out := make([]domain.Product, 0, limit)
	for i := 0; i < limit; i++ { out = append(out, bucket[i].P) }
	return out, nil
}

func sharedCount(a, b []string) int {
	m := map[string]struct{}{}
	for _, x := range a { m[x] = struct{}{} }
	c := 0
	for _, y := range b {
		if _, ok := m[y]; ok { c++ }
	}
	return c
}

func (s *ProductService) Related(id string, limit int) ([]domain.Product, error) {
	base, err := s.repo.GetByID(id)
	if err != nil { return nil, err }

	if len(base.RelatedIDs) > 0 {
		all, err := s.repo.All()
		if err != nil {
			if errors.Is(err, jsonstore.ErrNotFound) { return []domain.Product{}, nil }
			return nil, err
		}
		byID := map[string]domain.Product{}
		for _, p := range all { byID[p.ID] = p }

		var out []domain.Product
		for _, rid := range base.RelatedIDs {
			if p, ok := byID[rid]; ok && p.ID != base.ID {
				out = append(out, p)
			}
		}
		if limit > 0 && len(out) > limit { out = out[:limit] }
		return out, nil
	}

	all, err := s.repo.All()
	if err != nil {
		if errors.Is(err, jsonstore.ErrNotFound) { return []domain.Product{}, nil }
		return nil, err
	}
	var out []domain.Product
	for _, p := range all {
		if p.ID == base.ID { continue }
		if base.Category == "cellphone" && p.Category == "accessory" && sharedCount(base.Tags, p.Tags) > 0 {
			out = append(out, p)
		}
	}
	if limit > 0 && len(out) > limit { out = out[:limit] }
	return out, nil
}
