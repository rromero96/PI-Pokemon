package pokemon

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rromero96/roro-lib/cmd/rest"
)

func TestMakeGetPokemons_success(t *testing.T) {
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

	want := ErrPokemonNotFound
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

func TestMakeGetTypes_success(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusOK, []byte(`{}`)))

	_, got := MakeGetTypes(restGetFunc)

	assert.Nil(t, got)
}

func TestTestMakeGetTypes_success(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusOK, []byte(MockPokemonTypesAsJson())))
	getTypes, _ := MakeGetTypes(restGetFunc)
	ctx := context.Background()

	want := MockPokemonTypes()
	got, err := getTypes(ctx)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestTestMakeGetTypes_failsWithNotFound(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusNotFound, []byte{}))
	getTypes, _ := MakeGetTypes(restGetFunc)
	ctx := context.Background()

	want := ErrTypesNotFound
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}

func TestTestMakeGetTypes_failsWithUnmarshalError(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusOK, []byte("InvalidBody")))
	getTypes, _ := MakeGetTypes(restGetFunc)
	ctx := context.Background()

	want := ErrUnmarshalResponse
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}

func TestTestMakeGetTypes_failsWithInternalServerError(t *testing.T) {
	restGetFunc := rest.MakeGetFuncMock(rest.NewResponse(http.StatusInternalServerError, []byte("error")))
	getTypes, _ := MakeGetTypes(restGetFunc)
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
