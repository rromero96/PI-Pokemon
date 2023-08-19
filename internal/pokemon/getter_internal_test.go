package pokemon

import (
	"context"
	"net/http"
	"testing"

	"github.com/rromero96/roro-lib/rusty"
	"github.com/stretchr/testify/assert"
)

func TestMakeSearch_success(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "", Error: nil, StatusCode: http.StatusOK}

	_, got := MakeSearch(mockRequester)

	assert.Nil(t, got)
}

func TestSearch_success(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: MockPokemonAsJson(), Error: nil, StatusCode: http.StatusOK}
	search, _ := MakeSearch(mockRequester)
	ctx := context.Background()
	id := 1

	want := MockPokemon()
	got, err := search(ctx, &id, nil)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestSearch_failsWithNotFound(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "", Error: nil, StatusCode: http.StatusNotFound}
	search, _ := MakeSearch(mockRequester)
	ctx := context.Background()
	id := 1

	want := ErrPokemonNotFound
	_, got := search(ctx, &id, nil)

	assert.Equal(t, got, want)
}

func TestSearch_failsWithUnmarshalError(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: `{"error"}`, Error: nil, StatusCode: http.StatusOK}
	search, _ := MakeSearch(mockRequester)
	ctx := context.Background()
	id := 1

	want := ErrUnmarshalResponse
	_, got := search(ctx, &id, nil)

	assert.Equal(t, got, want)
}

func TestSearch_failsWithInternalServerError(t *testing.T) {
	mockRequester := rusty.RequesterResponseMock{Body: "error", Error: nil, StatusCode: http.StatusInternalServerError}
	search, _ := MakeSearch(mockRequester)
	ctx := context.Background()
	id := 1

	want := ErrCantGetPokemon
	_, got := search(ctx, &id, nil)

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
