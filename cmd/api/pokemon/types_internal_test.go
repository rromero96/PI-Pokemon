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
	types := MockTypes()
	types[0].ID = 1
	types[1].ID = 2

	want := types
	got := toTypesSlice(pokemonTypes.Types)

	assert.Equal(t, got, want)
}

func TestToPokemon_success(t *testing.T) {
	pokemon := pokemon.MockPokemon()

	want := MockPokemon()
	got := toPokemon(pokemon)

	assert.Equal(t, got, want)
}

func TestToType_success(t *testing.T) {
	types := pokemon.MockPokemon().Types

	want := MockPokemon().Types
	got := toType(types)

	assert.Equal(t, got, want)
}
