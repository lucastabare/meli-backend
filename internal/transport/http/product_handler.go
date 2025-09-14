package http

import (
	"errors"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
	"meli/internal/repository/jsonstore"
	"meli/internal/service"
	"meli/internal/util"
)

type ProductHandler struct{ svc *service.ProductService }
func NewProductHandler(s *service.ProductService) *ProductHandler { return &ProductHandler{s} }

func (h *ProductHandler) Routes(r chi.Router) {
	r.Get("/", h.list)
	r.Get("/{id}", h.get)
	r.Get("/{id}/seller", h.sellerRef)
	r.Get("/{id}/description", h.description)
	r.Get("/{id}/similar", h.similar)
	r.Get("/{id}/related", h.related)
}

func (h *ProductHandler) list(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	items, total, err := h.svc.List(service.ListParams{Query: q, Limit: limit, Offset: offset})
	if err != nil { util.Error(w, err); return }
	util.JSON(w, http.StatusOK, map[string]any{"total": total, "items": items})
}

func (h *ProductHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	item, err := h.svc.Get(id)
	if err != nil {
		if errors.Is(err, jsonstore.ErrNotFound) { util.Error(w, util.ErrNotFound); return }
		util.Error(w, err); return
	}
	util.JSON(w, http.StatusOK, item)
}

func (h *ProductHandler) sellerRef(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	item, err := h.svc.Get(id)
	if err != nil {
		if errors.Is(err, jsonstore.ErrNotFound) { util.Error(w, util.ErrNotFound); return }
		util.Error(w, err); return
	}
	util.JSON(w, http.StatusOK, map[string]string{"seller_id": item.SellerID})
}

func (h *ProductHandler) description(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	item, err := h.svc.Get(id)
	if err != nil {
		if errors.Is(err, jsonstore.ErrNotFound) { util.Error(w, util.ErrNotFound); return }
		util.Error(w, err); return
	}
	util.JSON(w, http.StatusOK, map[string]string{"id": item.ID, "description": item.Description})
}

func (h *ProductHandler) similar(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	items, err := h.svc.Similar(id, limit)
	if err != nil {
		if errors.Is(err, jsonstore.ErrNotFound) { util.Error(w, util.ErrNotFound); return }
		util.Error(w, err); return
	}
	util.JSON(w, http.StatusOK, items)
}

func (h *ProductHandler) related(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	items, err := h.svc.Related(id, limit)
	if err != nil {
		if errors.Is(err, jsonstore.ErrNotFound) { util.Error(w, util.ErrNotFound); return }
		util.Error(w, err); return
	}
	util.JSON(w, http.StatusOK, items)
}
