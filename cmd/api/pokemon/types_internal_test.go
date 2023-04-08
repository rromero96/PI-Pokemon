package pokemon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokemon_ToDTO_success(t *testing.T) {
	pokemon := MockPokemon()

	want := MockPokemonDTO()
	got := pokemon.toDTO()

	assert.Equal(t, got, want)
}

func TestToTypesDTO_succes(t *testing.T) {
	types := MockTypes()

	want := MockTypesDTO()
	got := toTypesDTO(types)

	assert.Equal(t, got, want)
}
