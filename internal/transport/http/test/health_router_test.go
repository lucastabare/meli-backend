package http_test

import (
	"meli/internal/repository/jsonstore"
	"meli/internal/service"
	transporthttp "meli/internal/transport/http"
	stdhttp "net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRouterFull_Health(t *testing.T) {
	td := t.TempDir()
	_ = os.WriteFile(filepath.Join(td, "products.json"), []byte(`[]`), 0o644)
	_ = os.WriteFile(filepath.Join(td, "sellers.json"), []byte(`[]`), 0o644)
	_ = os.WriteFile(filepath.Join(td, "payments.json"), []byte(`[]`), 0o644)
	_ = os.WriteFile(filepath.Join(td, "ads.json"), []byte(`[]`), 0o644)

	st := jsonstore.NewStore(td)
	productH := transporthttp.NewProductHandler(service.NewProductService(jsonstore.NewProductRepo(st)))
	sellerH := transporthttp.NewSellerHandler(service.NewSellerService(jsonstore.NewSellerRepo(st)))
	paymentH := transporthttp.NewPaymentHandler(service.NewPaymentService(jsonstore.NewPaymentRepo(st)))
	adsH := transporthttp.NewAdsHandler(service.NewAdsService(jsonstore.NewAdsRepo(st)))

	r := transporthttp.NewRouterFull(productH, sellerH, paymentH, adsH)

	req := httptest.NewRequest(stdhttp.MethodGet, "/api/health", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	require.Equal(t, stdhttp.StatusOK, rr.Code)
	require.Equal(t, "ok", rr.Body.String())
}
