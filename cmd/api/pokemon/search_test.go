package pokemon_test

import (
	"context"
	"testing"

	"github.com/rromero96/PI-Pokemon/cmd/api/pokemon"
	internalPokemon "github.com/rromero96/PI-Pokemon/internal/pokemon"
	"github.com/stretchr/testify/assert"
)

func TestMakeSearchTypes_success(t *testing.T) {
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(pokemon.MockTypes(), nil)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(nil)
	pokemonTypeSearch := internalPokemon.MockSearchTypes(internalPokemon.MockTypes(), nil)

	got := pokemon.MakeSearchTypes(mysqlSearchTypes, pokemonTypeSearch, mysqlCreateTypes)

	assert.NotNil(t, got)
}

func TestSearchTypes_success(t *testing.T) {
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(pokemon.MockTypes(), nil)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(nil)
	pokemonTypeSearch := internalPokemon.MockSearchTypes(internalPokemon.MockTypes(), nil)

	ctx := context.Background()
	searchTypes := pokemon.MakeSearchTypes(mysqlSearchTypes, pokemonTypeSearch, mysqlCreateTypes)

	want := pokemon.MockTypes()
	got, err := searchTypes(ctx)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestSearchTypes_failsWhenCantSearchTypes(t *testing.T) {
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(nil, pokemon.ErrCantSearchTypes)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(nil)
	pokemonTypeSearch := internalPokemon.MockSearchTypes(internalPokemon.MockTypes(), nil)

	ctx := context.Background()
	searchTypes := pokemon.MakeSearchTypes(mysqlSearchTypes, pokemonTypeSearch, mysqlCreateTypes)

	want := pokemon.ErrCantSearchTypes
	_, got := searchTypes(ctx)

	assert.Equal(t, got, want)
}

func TestSearchTypes_failsWhenCantSearchPokemonTypes(t *testing.T) {
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(nil, nil)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(nil)
	pokemonTypeSearch := internalPokemon.MockSearchTypes(internalPokemon.PokemonTypes{}, pokemon.ErrCantSearchPokemonTypes)

	ctx := context.Background()
	searchTypes := pokemon.MakeSearchTypes(mysqlSearchTypes, pokemonTypeSearch, mysqlCreateTypes)

	want := pokemon.ErrCantSearchPokemonTypes
	_, got := searchTypes(ctx)

	assert.Equal(t, got, want)
}

func TestSearchTypes_failsWhenCantSaveTypes(t *testing.T) {
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(nil, nil)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(pokemon.ErrCantSaveTypes)
	pokemonTypeSearch := internalPokemon.MockSearchTypes(internalPokemon.MockTypes(), nil)

	ctx := context.Background()
	searchTypes := pokemon.MakeSearchTypes(mysqlSearchTypes, pokemonTypeSearch, mysqlCreateTypes)

	want := pokemon.ErrCantSaveTypes
	_, got := searchTypes(ctx)

	assert.Equal(t, got, want)
}
