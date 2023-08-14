package pokemon

import (
	"context"
	"net/http"
	"testing"

	"github.com/rromero96/roro-lib/rusty"
	"github.com/stretchr/testify/assert"
)

func TestMakeSearchPokemons_success(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "", Error: nil, StatusCode: http.StatusOK}

	_, got := MakeSearchPokemon(mockRequester)

	assert.Nil(t, got)
}

func TestSearchPokemons_success(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: MockPokemonAsJson(), Error: nil, StatusCode: http.StatusOK}
	getPokemons, _ := MakeSearchPokemon(mockRequester)
	ctx := context.Background()
	id := 1

	want := MockPokemon()
	got, err := getPokemons(ctx, &id, nil)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestSearchPokemons_failsWithNotFound(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "", Error: nil, StatusCode: http.StatusNotFound}
	getPokemons, _ := MakeSearchPokemon(mockRequester)
	ctx := context.Background()
	id := 1

	want := ErrPokemonNotFound
	_, got := getPokemons(ctx, &id, nil)

	assert.Equal(t, got, want)
}

func TestSearchPokemons_failsWithUnmarshalError(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: `{"error"}`, Error: nil, StatusCode: http.StatusOK}
	getPokemons, _ := MakeSearchPokemon(mockRequester)
	ctx := context.Background()
	id := 1

	want := ErrUnmarshalResponse
	_, got := getPokemons(ctx, &id, nil)

	assert.Equal(t, got, want)
}

func TestSearchPokemons_failsWithInternalServerError(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "error", Error: nil, StatusCode: http.StatusInternalServerError}
	getPokemons, _ := MakeSearchPokemon(mockRequester)
	ctx := context.Background()
	id := 1

	want := ErrCantGetPokemon
	_, got := getPokemons(ctx, &id, nil)

	assert.Equal(t, got, want)
}

func TestMakeSearchTypes_success(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "", Error: nil, StatusCode: http.StatusOK}

	_, got := MakeSearchTypes(mockRequester)

	assert.Nil(t, got)
}

func TestSearchTypes_success(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: MockTypesAsJson(), Error: nil, StatusCode: http.StatusOK}
	getTypes, _ := MakeSearchTypes(mockRequester)
	ctx := context.Background()

	want := MockTypes()
	got, err := getTypes(ctx)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestSearchTypes_failsWithNotFound(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "", Error: nil, StatusCode: http.StatusNotFound}
	getTypes, _ := MakeSearchTypes(mockRequester)
	ctx := context.Background()

	want := ErrTypesNotFound
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}

func TestSearchTypes_failsWithUnmarshalError(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: `{"error"}`, Error: nil, StatusCode: http.StatusOK}
	getTypes, _ := MakeSearchTypes(mockRequester)
	ctx := context.Background()

	want := ErrUnmarshalResponse
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}

func TestSearchTypes_failsWithInternalServerError(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "error", Error: nil, StatusCode: http.StatusInternalServerError}
	getTypes, _ := MakeSearchTypes(mockRequester)
	ctx := context.Background()

	want := ErrCantGetTypes
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}
