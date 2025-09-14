package service_test

import (
	"meli/internal/repository/jsonstore"
	"meli/internal/service"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func writeFile(t *testing.T, dir, name, content string) {
	t.Helper()
	require.NoError(t, os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644))
}

func bootstrapStoreWithData(t *testing.T) *jsonstore.Store {
	t.Helper()
	td := t.TempDir()
	writeFile(t, td, "products.json", `[
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
	]`)
	writeFile(t, td, "sellers.json", `[{"id":"SELLER1","nickname":"Shop","city":"CABA","sales":100,"reputation":"green","rating_average":4.8}]`)
	writeFile(t, td, "payments.json", `[{"id":"visa","name":"Visa","type":"credit_card","installments":[1,3,6]}]`)
	return jsonstore.NewStore(td)
}

func TestProductService_List_Get_Similar_Related(t *testing.T) {
	st := bootstrapStoreWithData(t)
	repo := jsonstore.NewProductRepo(st)
	svc := service.NewProductService(repo)

	items, total, err := svc.List(service.ListParams{Query: "samsung", Limit: 10, Offset: 0})
	require.NoError(t, err)
	require.Greater(t, total, 0)
	require.NotEmpty(t, items)

	got, err := svc.Get("MLA123")
	require.NoError(t, err)
	require.Equal(t, "MLA123", got.ID)

	sim, err := svc.Similar("MLA123", 5)
	require.NoError(t, err)
	require.NotEmpty(t, sim)

	rel, err := svc.Related("MLA123", 5)
	require.NoError(t, err)
	foundACC := false
	for _, p := range rel {
		if p.ID == "ACC001" {
			foundACC = true
		}
	}
	require.True(t, foundACC)
}
