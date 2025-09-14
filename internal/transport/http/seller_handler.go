package http

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"meli/internal/service"
	"meli/internal/util"
)

type SellerHandler struct{ svc *service.SellerService }
func NewSellerHandler(s *service.SellerService) *SellerHandler { return &SellerHandler{s} }

func (h *SellerHandler) Routes(r chi.Router) {
	r.Get("/{id}", h.get)
}

func (h *SellerHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	s, err := h.svc.Get(id)
	if err != nil { util.Error(w, err); return }
	util.JSON(w, http.StatusOK, s)
}
