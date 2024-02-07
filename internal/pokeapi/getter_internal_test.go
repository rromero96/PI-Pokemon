package pokeapi

import (
	"context"
	"net/http"
	"testing"

	"github.com/rromero96/roro-lib/rusty"
	"github.com/stretchr/testify/assert"
)

func TestMakeGetByID_success(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "", Error: nil, StatusCode: http.StatusOK}

	_, got := MakeGetByID(mockRequester)

	assert.Nil(t, got)
}

func TestGetByID_success(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: MockPokemonAsJson(), Error: nil, StatusCode: http.StatusOK}
	getByID, _ := MakeGetByID(mockRequester)
	ctx := context.Background()
	id := 1

	want := MockPokemonDTO()
	got, err := getByID(ctx, id)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestGetByID_failsWithNotFound(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "", Error: nil, StatusCode: http.StatusNotFound}
	getByID, _ := MakeGetByID(mockRequester)
	ctx := context.Background()
	id := 1

	want := ErrPokemonNotFound
	_, got := getByID(ctx, id)

	assert.Equal(t, got, want)
}

func TestGetByID_failsWithUnmarshalError(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: `{"error"}`, Error: nil, StatusCode: http.StatusOK}
	getByID, _ := MakeGetByID(mockRequester)
	ctx := context.Background()
	id := 1

	want := ErrUnmarshalResponse
	_, got := getByID(ctx, id)

	assert.Equal(t, got, want)
}

func TestGetByID_failsWithInternalServerError(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "error", Error: nil, StatusCode: http.StatusInternalServerError}
	getByID, _ := MakeGetByID(mockRequester)
	ctx := context.Background()
	id := 1

	want := ErrCantGetPokemon
	_, got := getByID(ctx, id)

	assert.Equal(t, got, want)
}

func TestMakeGetTypes_success(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "", Error: nil, StatusCode: http.StatusOK}

	_, got := MakeGetTypes(mockRequester)

	assert.Nil(t, got)
}

func TestGetTypes_success(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: MockTypesAsJson(), Error: nil, StatusCode: http.StatusOK}
	getTypes, _ := MakeGetTypes(mockRequester)
	ctx := context.Background()

	want := MockPokemonTypesDTO()
	got, err := getTypes(ctx)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestGetTypes_failsWithNotFound(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "", Error: nil, StatusCode: http.StatusNotFound}
	getTypes, _ := MakeGetTypes(mockRequester)
	ctx := context.Background()

	want := ErrTypesNotFound
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}

func TestGetTypes_failsWithUnmarshalError(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: `{"error"}`, Error: nil, StatusCode: http.StatusOK}
	getTypes, _ := MakeGetTypes(mockRequester)
	ctx := context.Background()

	want := ErrUnmarshalResponse
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}

func TestGetTypes_failsWithInternalServerError(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "error", Error: nil, StatusCode: http.StatusInternalServerError}
	getTypes, _ := MakeGetTypes(mockRequester)
	ctx := context.Background()

	want := ErrCantGetTypes
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}
