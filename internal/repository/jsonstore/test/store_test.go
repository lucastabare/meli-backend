package jsonstore_test

import (
	"meli/internal/repository/jsonstore"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductRepo_All_FileNotFound(t *testing.T) {
	st := jsonstore.NewStore(t.TempDir())
	repo := jsonstore.NewProductRepo(st)
	_, err := repo.All()
	require.Error(t, err)
	require.ErrorIs(t, err, jsonstore.ErrNotFound)
}
