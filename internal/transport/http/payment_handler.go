package http

import (
	"net/http"
	"github.com/go-chi/chi/v5"    
	"meli/internal/service"
	"meli/internal/util"
)

type PaymentHandler struct{ svc *service.PaymentService }

func NewPaymentHandler(s *service.PaymentService) *PaymentHandler { return &PaymentHandler{s} }

func (h *PaymentHandler) Routes(r chi.Router) { 
	r.Get("/methods", h.methods)
}

func (h *PaymentHandler) methods(w http.ResponseWriter, r *http.Request) {
	m, err := h.svc.Methods()
	if err != nil { util.Error(w, err); return }
	util.JSON(w, http.StatusOK, m)
}
