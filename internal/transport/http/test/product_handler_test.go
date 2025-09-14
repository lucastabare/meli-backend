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

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func setupProducts(t *testing.T) *transporthttp.ProductHandler {
	td := t.TempDir()
	_ = os.WriteFile(filepath.Join(td, "products.json"), []byte(`[
	  {"id":"MLA123","title":"Samsung Galaxy A55","price":439,"currency":"USD","condition":"new",
	   "pictures":[{"id":"p1","url":"u"}],"thumbnail":"t","permalink":"p","seller_id":"SELLER1",
	   "category":"cellphone","brand":"Samsung","tags":["samsung","galaxy","a55"],"related_ids":["ACC001"],
	   "shipping":{"free_shipping":true,"mode":"me2"},"stock":{"available":10,"sold":100},
	   "attributes":["a"],"description":"desc","rating_avg":4.6,"ratings":[]},
	  {"id":"MLA124","title":"Samsung Galaxy A35","price":329,"currency":"USD","condition":"new",
	   "pictures":[{"id":"p1","url":"u"}],"thumbnail":"t","permalink":"p","seller_id":"SELLER1",
	   "category":"cellphone","brand":"Samsung","tags":["samsung","a35"],
	   "shipping":{"free_shipping":true,"mode":"me2"},"stock":{"available":20,"sold":90},
	   "attributes":["a"],"description":"desc","rating_avg":4.4,"ratings":[]},
	  {"id":"ACC001","title":"Funda Samsung A55","price":12.9,"currency":"USD","condition":"new",
	   "pictures":[{"id":"p1","url":"u"}],"thumbnail":"t","permalink":"p","seller_id":"SELLER1",
	   "category":"accessory","brand":"Generic","tags":["samsung","a55","case"],
	   "shipping":{"free_shipping":false,"mode":"me2"},"stock":{"available":120,"sold":400},
	   "attributes":["a"],"description":"desc","rating_avg":4.7,"ratings":[]}
	]`), 0o644)
	_ = os.WriteFile(filepath.Join(td, "sellers.json"), []byte(`[{"id":"SELLER1","nickname":"Shop","city":"CABA","sales":100,"reputation":"green","rating_average":4.8}]`), 0o644)
	_ = os.WriteFile(filepath.Join(td, "payments.json"), []byte(`[{"id":"visa","name":"Visa","type":"credit_card","installments":[1,3,6]}]`), 0o644)

	store := jsonstore.NewStore(td)
	productSvc := service.NewProductService(jsonstore.NewProductRepo(store))
	return transporthttp.NewProductHandler(productSvc)
}

func TestProductHandler_List_Get_Similar_Related(t *testing.T) {
	h := setupProducts(t)
	r := chi.NewRouter()
	h.Routes(r)

	// list
	req := httptest.NewRequest(stdhttp.MethodGet, "/?q=samsung&limit=5", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	require.Equal(t, stdhttp.StatusOK, rr.Code)
	require.Contains(t, rr.Body.String(), `"items"`)

	// detail
	req = httptest.NewRequest(stdhttp.MethodGet, "/MLA123", nil)
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	require.Equal(t, stdhttp.StatusOK, rr.Code)
	require.Contains(t, rr.Body.String(), `"MLA123"`)

	// similar
	req = httptest.NewRequest(stdhttp.MethodGet, "/MLA123/similar?limit=2", nil)
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	require.Equal(t, stdhttp.StatusOK, rr.Code)

	// related
	req = httptest.NewRequest(stdhttp.MethodGet, "/MLA123/related?limit=2", nil)
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	require.Equal(t, stdhttp.StatusOK, rr.Code)
	require.Contains(t, rr.Body.String(), "ACC001")
}
