package http_test

import (
	stdhttp "net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"meli/internal/repository/jsonstore"
	"meli/internal/service"
	transporthttp "meli/internal/transport/http"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestSellerHandler_Get(t *testing.T) {
	td := t.TempDir()
	_ = os.WriteFile(filepath.Join(td, "sellers.json"), []byte(`[{"id":"SELLER1","nickname":"Shop","city":"CABA","sales":100,"reputation":"green","rating_average":4.8}]`), 0o644)
	store := jsonstore.NewStore(td)
	svc := service.NewSellerService(jsonstore.NewSellerRepo(store))
	h := transporthttp.NewSellerHandler(svc)

	r := chi.NewRouter()
	h.Routes(r)

	req := httptest.NewRequest(stdhttp.MethodGet, "/SELLER1", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	require.Equal(t, stdhttp.StatusOK, rr.Code)
	require.Contains(t, rr.Body.String(), `"SELLER1"`)
}

func TestPaymentHandler_Methods(t *testing.T) {
	td := t.TempDir()
	_ = os.WriteFile(filepath.Join(td, "payments.json"), []byte(`[{"id":"visa","name":"Visa","type":"credit_card","installments":[1,3,6]}]`), 0o644)
	store := jsonstore.NewStore(td)
	svc := service.NewPaymentService(jsonstore.NewPaymentRepo(store))
	h := transporthttp.NewPaymentHandler(svc)

	r := chi.NewRouter()
	h.Routes(r)

	req := httptest.NewRequest(stdhttp.MethodGet, "/methods", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	require.Equal(t, stdhttp.StatusOK, rr.Code)
	require.Contains(t, rr.Body.String(), `"visa"`)
}
