package http

import (
	"net/http"

	"meli/internal/service"
	"meli/internal/util"

	"github.com/go-chi/chi/v5"
)

type AdsHandler struct{ svc *service.AdsService }

func NewAdsHandler(s *service.AdsService) *AdsHandler { return &AdsHandler{s} }

func (h *AdsHandler) Routes(r chi.Router) {
	r.Get("/", h.list)
}

func (h *AdsHandler) list(w http.ResponseWriter, r *http.Request) {
	ads, err := h.svc.List()
	if err != nil {
		util.Error(w, err)
		return
	}
	util.JSON(w, http.StatusOK, ads)
}
