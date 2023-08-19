package pokemon_test

import (
	"context"
	"errors"
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
	err := errors.New("error")
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(nil, err)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(nil)
	pokemonTypeSearch := internalPokemon.MockSearchTypes(internalPokemon.MockTypes(), nil)

	ctx := context.Background()
	searchTypes := pokemon.MakeSearchTypes(mysqlSearchTypes, pokemonTypeSearch, mysqlCreateTypes)

	want := pokemon.ErrCantSearchTypes
	_, got := searchTypes(ctx)

	assert.Equal(t, got, want)
}

func TestSearchTypes_failsWhenCantSearchPokemonTypes(t *testing.T) {
	err := errors.New("error")
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(nil, nil)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(nil)
	pokemonTypeSearch := internalPokemon.MockSearchTypes(internalPokemon.PokemonTypes{}, err)

	ctx := context.Background()
	searchTypes := pokemon.MakeSearchTypes(mysqlSearchTypes, pokemonTypeSearch, mysqlCreateTypes)

	want := pokemon.ErrCantSearchPokemonTypes
	_, got := searchTypes(ctx)

	assert.Equal(t, got, want)
}

func TestSearchTypes_failsWhenCantSaveTypes(t *testing.T) {
	err := errors.New("error")
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(nil, nil)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(err)
	pokemonTypeSearch := internalPokemon.MockSearchTypes(internalPokemon.MockTypes(), nil)

	ctx := context.Background()
	searchTypes := pokemon.MakeSearchTypes(mysqlSearchTypes, pokemonTypeSearch, mysqlCreateTypes)

	want := pokemon.ErrCantSaveTypes
	_, got := searchTypes(ctx)

	assert.Equal(t, got, want)
}

func TestMakeSearchByID_success(t *testing.T) {
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.MockPokemon(), nil)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokemonSearch := internalPokemon.MockSearch(internalPokemon.Pokemon{}, nil)

	got := pokemon.MakeSearchByID(mysqlSearchByID, pokemonSearch, mysqlCreate)

	assert.NotNil(t, got)
}

func TestSearchByID_success(t *testing.T) {
	ID := 1
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.MockPokemon(), nil)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokemonSearch := internalPokemon.MockSearch(internalPokemon.Pokemon{}, nil)

	ctx := context.Background()
	searchByID := pokemon.MakeSearchByID(mysqlSearchByID, pokemonSearch, mysqlCreate)

	want := pokemon.MockPokemon()
	got, err := searchByID(ctx, ID)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestSearchByID_successWhenPokemonIsNotInTheDatabase(t *testing.T) {
	ID := 1
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.Pokemon{}, nil)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokemonSearch := internalPokemon.MockSearch(internalPokemon.MockPokemon(), nil)

	ctx := context.Background()
	searchByID := pokemon.MakeSearchByID(mysqlSearchByID, pokemonSearch, mysqlCreate)

	want := pokemon.MockPokemon()
	got, err := searchByID(ctx, ID)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestSearchByID_failsWhenCantSearchByID(t *testing.T) {
	ID := 1
	err := errors.New("error")
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.Pokemon{}, err)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokemonSearch := internalPokemon.MockSearch(internalPokemon.MockPokemon(), nil)

	ctx := context.Background()
	searchByID := pokemon.MakeSearchByID(mysqlSearchByID, pokemonSearch, mysqlCreate)

	want := pokemon.ErrCantSearchPokemon
	_, got := searchByID(ctx, ID)

	assert.Equal(t, got, want)
}

func TestSearchByID_failsWhenCantSearchPokemonInPokeapi(t *testing.T) {
	ID := 1
	err := errors.New("error")
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.Pokemon{}, nil)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokemonSearch := internalPokemon.MockSearch(internalPokemon.Pokemon{}, err)

	ctx := context.Background()
	searchByID := pokemon.MakeSearchByID(mysqlSearchByID, pokemonSearch, mysqlCreate)

	want := pokemon.ErrCantSearchPokemonApi
	_, got := searchByID(ctx, ID)

	assert.Equal(t, got, want)
}

func TestSearchByID_failsWhenCantSavePokemon(t *testing.T) {
	ID := 1
	err := errors.New("error")
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.Pokemon{}, nil)
	mysqlCreate := pokemon.MockMySQLCreate(err)
	pokemonSearch := internalPokemon.MockSearch(internalPokemon.MockPokemon(), nil)

	ctx := context.Background()
	searchByID := pokemon.MakeSearchByID(mysqlSearchByID, pokemonSearch, mysqlCreate)

	want := pokemon.ErrCantCreatePokemon
	_, got := searchByID(ctx, ID)

	assert.Equal(t, got, want)
}
