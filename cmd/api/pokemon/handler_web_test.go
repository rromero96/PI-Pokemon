package pokemon_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rromero96/roro-lib/web"
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

func TestHTTPHandler_CreateV1_failsWhenBodyIsInvalid(t *testing.T) {
	createPokemon := pokemon.MockMySQLCreate(nil)
	createV1 := pokemon.CreateV1(createPokemon)
	bodyJSON := pokemon.InvalidBody

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", strings.NewReader(bodyJSON))

	want := web.NewError(http.StatusBadRequest, pokemon.InvalidBody)
	got := createV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHandler_CreateV1_failsWithBadRequest(t *testing.T) {
	createPokemon := pokemon.MockMySQLCreate(pokemon.ErrCantRunQuery)
	createV1 := pokemon.CreateV1(createPokemon)
	bodyJSON := pokemon.MockPokemonAsJson()

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", strings.NewReader(bodyJSON))

	want := web.NewError(http.StatusBadRequest, pokemon.InvalidPokemon)
	got := createV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHandler_CreateV1_failsWithInternalServerError(t *testing.T) {
	createPokemon := pokemon.MockMySQLCreate(pokemon.ErrCantPrepareStatement)
	createV1 := pokemon.CreateV1(createPokemon)
	bodyJSON := pokemon.MockPokemonAsJson()

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", strings.NewReader(bodyJSON))

	want := web.NewError(http.StatusInternalServerError, pokemon.CantCreatePokemon)
	got := createV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHanldler_SearchTypesV1_success(t *testing.T) {
	searchTypes := pokemon.MockSearchTypes(pokemon.MockTypes(), nil)
	searchTypesV1 := pokemon.SearchTypesV1(searchTypes)

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", nil)

	got := searchTypesV1(w, r)

	assert.Nil(t, got)
}

func TestHTTPHanldler_SearchTypesV1_failsWithInternalServerError(t *testing.T) {
	searchTypes := pokemon.MockSearchTypes(pokemon.MockTypes(), pokemon.ErrCantSearchTypes)
	searchTypesV1 := pokemon.SearchTypesV1(searchTypes)

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", nil)

	want := web.NewError(http.StatusInternalServerError, pokemon.CantGetTypes)
	got := searchTypesV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHandler_SearchByIDV1_success(t *testing.T) {
	searchByID := pokemon.MockSearchByID(pokemon.MockPokemon(), nil)
	searchByIDV1 := pokemon.SearchByIDV1(searchByID)

	w := httptest.NewRecorder()
	ctx := web.WithParams(context.Background(), web.URIParams{
		pokemon.ParamPokemonID: "1",
	})
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/test", nil)

	got := searchByIDV1(w, r)

	assert.Nil(t, got)
}

func TestHTTPHandler_SearchByIDV1_failsWithBadRequest(t *testing.T) {
	searchByID := pokemon.MockSearchByID(pokemon.MockPokemon(), nil)
	searchByIDV1 := pokemon.SearchByIDV1(searchByID)

	w := httptest.NewRecorder()
	ctx := web.WithParams(context.Background(), web.URIParams{
		pokemon.ParamPokemonID: "invalid",
	})
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/test", nil)

	want := web.NewError(http.StatusBadRequest, pokemon.InvalidID)
	got := searchByIDV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHandler_SearchByIDV1_failsWithNotFound(t *testing.T) {
	searchByID := pokemon.MockSearchByID(pokemon.Pokemon{}, pokemon.ErrPokemonNotFound)
	searchByIDV1 := pokemon.SearchByIDV1(searchByID)

	w := httptest.NewRecorder()
	ctx := web.WithParams(context.Background(), web.URIParams{
		pokemon.ParamPokemonID: "1",
	})
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/test", nil)

	want := web.NewError(http.StatusNotFound, pokemon.NotFound)
	got := searchByIDV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHandler_SearchByIDV1_failsWithInternalServerError(t *testing.T) {
	searchByID := pokemon.MockSearchByID(pokemon.Pokemon{}, pokemon.ErrCantSearchPokemon)
	searchByIDV1 := pokemon.SearchByIDV1(searchByID)

	w := httptest.NewRecorder()
	ctx := web.WithParams(context.Background(), web.URIParams{
		pokemon.ParamPokemonID: "1",
	})
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/test", nil)

	want := web.NewError(http.StatusInternalServerError, pokemon.CantSearchPokemon)
	got := searchByIDV1(w, r)

	assert.Equal(t, got, want)
}
