package pokemon

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rromero96/roro-lib/cmd/rest"
)

func TestMakeSearchPokemons_success(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusOK, []byte(`{}`)))

	_, got := MakeSearchPokemon(restGetFunc)

	assert.Nil(t, got)
}

func TestSearchPokemons_success(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusOK, []byte(MockPokemonAsJson())))
	getPokemons, _ := MakeSearchPokemon(restGetFunc)
	ctx := context.Background()
	id := 1

	want := MockPokemon()
	got, err := getPokemons(ctx, id)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestSearchPokemons_failsWithNotFound(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusNotFound, []byte{}))
	getPokemons, _ := MakeSearchPokemon(restGetFunc)
	ctx := context.Background()
	id := 1

	want := ErrPokemonNotFound
	_, got := getPokemons(ctx, id)

	assert.Equal(t, got, want)
}

func TestSearchPokemons_failsWithUnmarshalError(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusOK, []byte("InvalidBody")))
	getPokemons, _ := MakeSearchPokemon(restGetFunc)
	ctx := context.Background()
	id := 1

	want := ErrUnmarshalResponse
	_, got := getPokemons(ctx, id)

	assert.Equal(t, got, want)
}

func TestSearchPokemons_failsWithInternalServerError(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusInternalServerError, []byte("error")))
	getPokemons, _ := MakeSearchPokemon(restGetFunc)
	ctx := context.Background()
	id := 1
	requestURL := "https://pokeapi.co/api/v2/pokemon/1"

	want := rest.RequestError{
		Method:          http.MethodGet,
		URL:             requestURL,
		StatusCode:      http.StatusInternalServerError,
		ResponsePayload: "error",
	}
	_, got := getPokemons(ctx, id)

	assert.Equal(t, got, want)
}

func TestMakeSearchTypes_success(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusOK, []byte(`{}`)))

	_, got := MakeSearchTypes(restGetFunc)

	assert.Nil(t, got)
}

func TestTestMakeSearchTypes_success(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusOK, []byte(MockPokemonTypesAsJson())))
	getTypes, _ := MakeSearchTypes(restGetFunc)
	ctx := context.Background()

	want := MockPokemonTypes()
	got, err := getTypes(ctx)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestTestMakeSearchTypes_failsWithNotFound(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusNotFound, []byte{}))
	getTypes, _ := MakeSearchTypes(restGetFunc)
	ctx := context.Background()

	want := ErrTypesNotFound
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}

func TestTestMakeSearchTypes_failsWithUnmarshalError(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusOK, []byte("InvalidBody")))
	getTypes, _ := MakeSearchTypes(restGetFunc)
	ctx := context.Background()

	want := ErrUnmarshalResponse
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}

func TestTestMakeSearchTypes_failsWithInternalServerError(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusInternalServerError, []byte("error")))
	getTypes, _ := MakeSearchTypes(restGetFunc)
	ctx := context.Background()
	requestURL := "https://pokeapi.co/api/v2/type"

	want := rest.RequestError{
		Method:          http.MethodGet,
		URL:             requestURL,
		StatusCode:      http.StatusInternalServerError,
		ResponsePayload: "error",
	}
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}
