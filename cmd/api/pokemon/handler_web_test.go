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
	createPokemon := pokemon.MockCreate(nil)
	createV1 := pokemon.CreateV1(createPokemon)
	bodyJSON := pokemon.MockPokemonAsJson()

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", strings.NewReader(bodyJSON))

	got := createV1(w, r)

	assert.Nil(t, got)
}

func TestHTTPHandler_CreateV1_failsWhenBodyIsInvalid(t *testing.T) {
	createPokemon := pokemon.MockCreate(nil)
	createV1 := pokemon.CreateV1(createPokemon)
	bodyJSON := pokemon.InvalidBody

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", strings.NewReader(bodyJSON))

	want := web.NewError(http.StatusBadRequest, pokemon.InvalidBody)
	got := createV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHandler_CreateV1_failsWithBadRequest(t *testing.T) {
	createPokemon := pokemon.MockCreate(pokemon.ErrInvalidPokemon)
	createV1 := pokemon.CreateV1(createPokemon)
	bodyJSON := pokemon.MockPokemonAsJson()

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", strings.NewReader(bodyJSON))

	want := web.NewError(http.StatusBadRequest, pokemon.InvalidPokemon)
	got := createV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHandler_CreateV1_failsWithInternalServerError(t *testing.T) {
	createPokemon := pokemon.MockCreate(pokemon.ErrCantPrepareStatement)
	createV1 := pokemon.CreateV1(createPokemon)
	bodyJSON := pokemon.MockPokemonAsJson()

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", strings.NewReader(bodyJSON))

	want := web.NewError(http.StatusInternalServerError, pokemon.CantCreatePokemon)
	got := createV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHanldler_GetTypesV1_success(t *testing.T) {
	getTypes := pokemon.MockGetTypes(pokemon.MockTypes(), nil)
	getTypesV1 := pokemon.GetTypesV1(getTypes)

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", nil)

	got := getTypesV1(w, r)

	assert.Nil(t, got)
}

func TestHTTPHanldler_GetTypesV1_failsWithInternalServerError(t *testing.T) {
	getTypes := pokemon.MockGetTypes(pokemon.MockTypes(), pokemon.ErrCantGetTypes)
	getTypesV1 := pokemon.GetTypesV1(getTypes)

	ctx, w := context.Background(), httptest.NewRecorder()
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/test", nil)

	want := web.NewError(http.StatusInternalServerError, pokemon.CantGetTypes)
	got := getTypesV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHandler_GetByIDV1_success(t *testing.T) {
	getByID := pokemon.MockGetByID(pokemon.MockPokemon(), nil)
	getByIDV1 := pokemon.GetByIDV1(getByID)

	w := httptest.NewRecorder()
	ctx := web.WithParams(context.Background(), web.URIParams{
		pokemon.ParamPokemonID: "1",
	})
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/test", nil)

	got := getByIDV1(w, r)

	assert.Nil(t, got)
}

func TestHTTPHandler_GetByIDV1_failsWithBadRequest(t *testing.T) {
	getByID := pokemon.MockGetByID(pokemon.MockPokemon(), nil)
	getByIDV1 := pokemon.GetByIDV1(getByID)

	w := httptest.NewRecorder()
	ctx := web.WithParams(context.Background(), web.URIParams{
		pokemon.ParamPokemonID: "invalid",
	})
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/test", nil)

	want := web.NewError(http.StatusBadRequest, pokemon.InvalidID)
	got := getByIDV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHandler_GetByIDV1_failsWithNotFound(t *testing.T) {
	getByID := pokemon.MockGetByID(pokemon.Pokemon{}, pokemon.ErrPokemonNotFound)
	getByIDV1 := pokemon.GetByIDV1(getByID)

	w := httptest.NewRecorder()
	ctx := web.WithParams(context.Background(), web.URIParams{
		pokemon.ParamPokemonID: "1",
	})
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/test", nil)

	want := web.NewError(http.StatusNotFound, pokemon.NotFound)
	got := getByIDV1(w, r)

	assert.Equal(t, got, want)
}

func TestHTTPHandler_GetByIDV1_failsWithInternalServerError(t *testing.T) {
	getByID := pokemon.MockGetByID(pokemon.Pokemon{}, pokemon.ErrCantGetPokemon)
	getByIDV1 := pokemon.GetByIDV1(getByID)

	w := httptest.NewRecorder()
	ctx := web.WithParams(context.Background(), web.URIParams{
		pokemon.ParamPokemonID: "1",
	})
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/test", nil)

	want := web.NewError(http.StatusInternalServerError, pokemon.CantSearchPokemon)
	got := getByIDV1(w, r)

	assert.Equal(t, got, want)
}
