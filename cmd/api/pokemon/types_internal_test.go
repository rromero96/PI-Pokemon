package pokemon

import (
	"testing"

	"github.com/rromero96/PI-Pokemon/internal/pokemon"
	"github.com/stretchr/testify/assert"
)

func TestPokemon_ToDTO_success(t *testing.T) {
	pokemon := MockPokemon()

	want := MockPokemonDTO()
	got := pokemon.toDTO()

	assert.Equal(t, got, want)
}

func TestToTypesDTO_success(t *testing.T) {
	types := MockTypes()

	want := MockTypesDTO()
	got := toTypesDTO(types)

	assert.Equal(t, got, want)
}

func TestToTypesSlice_success(t *testing.T) {
	pokemonTypes := pokemon.MockTypes()

	want := MockTypes()
	got := toTypesSlice(pokemonTypes.Types)

	assert.Equal(t, got, want)
}
