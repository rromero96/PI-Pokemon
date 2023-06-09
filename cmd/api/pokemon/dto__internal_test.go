package pokemon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokemon_Validate_success(t *testing.T) {
	pokemon := MockPokemonDTO()

	got := pokemon.validate()

	assert.Nil(t, got)
}

func TestPokemon_Validate_failsWithInvalidBody(t *testing.T) {
	pokemon := MockPokemonDTO()
	pokemon.Types = nil

	want := ErrInvalidBody
	got := pokemon.validate()

	assert.Equal(t, got, want)
}

func TestTypes_Validate_success(t *testing.T) {
	types := MockTypesDTO()

	got := types[0].validate()

	assert.Nil(t, got)
}

func TestTypes_Validate_failsWithInvalidBody(t *testing.T) {
	types := MockTypesDTO()
	types[0].Name = ""

	want := ErrInvalidBody
	got := types[0].validate()

	assert.Equal(t, got, want)
}

func TestPokemon_ToDomain_success(t *testing.T) {
	pokemon := MockPokemonDTO()

	want := MockPokemon()
	got := pokemon.toDomain()

	assert.Equal(t, got, want)
}

func TestToTypes_success(t *testing.T) {
	types := MockTypesDTO()

	want := MockTypes()
	got := toTypes(types)

	assert.Equal(t, got, want)
}
