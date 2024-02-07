package pokemon

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rromero96/PI-Pokemon/internal/pokeapi"
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
	pokemonTypes := pokeapi.MockPokemonTypesDTO()
	types := MockTypes()
	types[0].ID = 1
	types[1].ID = 2

	want := types
	got := toTypesSlice(pokemonTypes.Types)

	assert.Equal(t, got, want)
}

func TestToPokemon_success(t *testing.T) {
	pokemon := pokeapi.MockPokemonDTO()

	want := MockPokemon()
	got := toPokemon(pokemon)

	assert.Equal(t, got, want)
}

func TestToType_success(t *testing.T) {
	types := pokeapi.MockPokemonDTO().Types

	want := MockPokemon().Types
	got := toType(types)

	assert.Equal(t, got, want)
}
