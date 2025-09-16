package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouterFull(
	product *ProductHandler,
	seller *SellerHandler,
	payment *PaymentHandler,
) http.Handler {
	r := chi.NewRouter()
	r.Use(Logging())
	r.Use(CORSManual())

	r.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	r.Route("/api/v1", func(v chi.Router) {
		v.Route("/products", product.Routes)
		v.Route("/sellers", seller.Routes)
		v.Route("/payments", payment.Routes)
	})

	return r
}
