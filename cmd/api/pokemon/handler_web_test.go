package pokemon_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rromero96/roro-lib/cmd/web"
	"github.com/stretchr/testify/assert"

	"github.com/rromero96/PI-Pokemon/cmd/api/pokemon"
)

func TestHTTPHandler_CreateV1_success(t *testing.T) {
	createPokemon := pokemon.MockMySQLCreate(nil)
	createV1 := pokemon.CreateV1(createPokemon)
	bodyJSON := pokemon.MockPokemonAsJson()

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", strings.NewReader(bodyJSON))

	got := createV1(w, r)

	assert.Nil(t, got)
}

func TestHTTPHandler__failsWhenBodyIsInvalid(t *testing.T) {
	createPokemon := pokemon.MockMySQLCreate(nil)
	createV1 := pokemon.CreateV1(createPokemon)
	bodyJSON := pokemon.InvalidBody

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", strings.NewReader(bodyJSON))

	want := web.NewError(http.StatusBadRequest, pokemon.InvalidBody)
	got := createV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHandler__failsWithBadRequest(t *testing.T) {
	createPokemon := pokemon.MockMySQLCreate(pokemon.ErrCantRunQuery)
	createV1 := pokemon.CreateV1(createPokemon)
	bodyJSON := pokemon.MockPokemonAsJson()

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", strings.NewReader(bodyJSON))

	want := web.NewError(http.StatusBadRequest, pokemon.InvalidPokemon)
	got := createV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHandler__failsWithInternalServerError(t *testing.T) {
	createPokemon := pokemon.MockMySQLCreate(pokemon.ErrCantPrepareStatement)
	createV1 := pokemon.CreateV1(createPokemon)
	bodyJSON := pokemon.MockPokemonAsJson()

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", strings.NewReader(bodyJSON))

	want := web.NewError(http.StatusInternalServerError, pokemon.CantCreatePokemon)
	got := createV1(w, r)

	assert.Equal(t, got, want)
}
