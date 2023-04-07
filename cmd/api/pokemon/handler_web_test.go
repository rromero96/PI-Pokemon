package pokemon_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rromero96/PI-Pokemon/cmd/api/pokemon"
	"github.com/stretchr/testify/assert"
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

}

func TestHTTPHandler__failsWithBadRequest(t *testing.T) {

}

func TestHTTPHandler__failsWithInternalServerError(t *testing.T) {

}
