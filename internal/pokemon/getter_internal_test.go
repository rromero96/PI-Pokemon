package pokemon

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rromero96/roro-lib/cmd/rest"
)

func TestMakeGetPokemons(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusOK, []byte(`{}`)))

	_, got := MakeGetPokemon(restGetFunc)

	assert.Nil(t, got)
}

func TestGetPokemons_success(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusOK, []byte(MockPokemonAsJson())))
	getPokemons, _ := MakeGetPokemon(restGetFunc)
	ctx := context.Background()
	id := 1

	want := MockPokemon()
	got, err := getPokemons(ctx, id)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestGetPokemons_failsWithNotFound(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusNotFound, []byte{}))
	getPokemons, _ := MakeGetPokemon(restGetFunc)
	ctx := context.Background()
	id := 1

	want := ErrNotFound
	_, got := getPokemons(ctx, id)

	assert.Equal(t, got, want)
}

func TestGetPokemons_failsWithUnmarshalError(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusOK, []byte("InvalidBody")))
	getPokemons, _ := MakeGetPokemon(restGetFunc)
	ctx := context.Background()
	id := 1

	want := ErrUnmarshalResponse
	_, got := getPokemons(ctx, id)

	assert.Equal(t, got, want)
}

func TestGetPokemons_failsWithInternalServerError(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusInternalServerError, []byte("error")))
	getPokemons, _ := MakeGetPokemon(restGetFunc)
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
