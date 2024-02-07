package pokemon_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rromero96/PI-Pokemon/cmd/api/pokemon"
	"github.com/rromero96/PI-Pokemon/internal/pokeapi"
)

func TestMakeSearchTypes_success(t *testing.T) {
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(pokemon.MockTypes(), nil)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(nil)
	pokeapiGetTypes := pokeapi.MockGetTypes(pokeapi.MockTypes(), nil)

	got := pokemon.MakeSearchTypes(mysqlSearchTypes, pokeapiGetTypes, mysqlCreateTypes)

	assert.NotNil(t, got)
}

func TestSearchTypes_success(t *testing.T) {
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(pokemon.MockTypes(), nil)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(nil)
	pokeapiGetTypes := pokeapi.MockGetTypes(pokeapi.MockTypes(), nil)

	ctx := context.Background()
	searchTypes := pokemon.MakeSearchTypes(mysqlSearchTypes, pokeapiGetTypes, mysqlCreateTypes)

	want := pokemon.MockTypes()
	got, err := searchTypes(ctx)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestSearchTypes_failsWhenCantSearchTypes(t *testing.T) {
	err := errors.New("error")
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(nil, err)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(nil)
	pokeapiGetTypes := pokeapi.MockGetTypes(pokeapi.MockTypes(), nil)

	ctx := context.Background()
	searchTypes := pokemon.MakeSearchTypes(mysqlSearchTypes, pokeapiGetTypes, mysqlCreateTypes)

	want := pokemon.ErrCantSearchTypes
	_, got := searchTypes(ctx)

	assert.Equal(t, got, want)
}

func TestSearchTypes_failsWhenCantGetPokemonTypes(t *testing.T) {
	err := errors.New("error")
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(nil, nil)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(nil)
	pokeapiGetTypes := pokeapi.MockGetTypes(pokeapi.PokemonTypes{}, err)

	ctx := context.Background()
	searchTypes := pokemon.MakeSearchTypes(mysqlSearchTypes, pokeapiGetTypes, mysqlCreateTypes)

	want := pokemon.ErrCantGetPokemonTypes
	_, got := searchTypes(ctx)

	assert.Equal(t, got, want)
}

func TestSearchTypes_failsWhenCantSaveTypes(t *testing.T) {
	err := errors.New("error")
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(nil, nil)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(err)
	pokeapiGetTypes := pokeapi.MockGetTypes(pokeapi.MockTypes(), nil)

	ctx := context.Background()
	searchTypes := pokemon.MakeSearchTypes(mysqlSearchTypes, pokeapiGetTypes, mysqlCreateTypes)

	want := pokemon.ErrCantSaveTypes
	_, got := searchTypes(ctx)

	assert.Equal(t, got, want)
}

func TestMakeSearchByID_success(t *testing.T) {
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.MockPokemon(), nil)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokeapiGetByID := pokeapi.MockGetByID(pokeapi.Pokemon{}, nil)

	got := pokemon.MakeSearchByID(mysqlSearchByID, pokeapiGetByID, mysqlCreate)

	assert.NotNil(t, got)
}

func TestSearchByID_success(t *testing.T) {
	ID := 1
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.MockPokemon(), nil)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokeapiGetByID := pokeapi.MockGetByID(pokeapi.Pokemon{}, nil)

	ctx := context.Background()
	searchByID := pokemon.MakeSearchByID(mysqlSearchByID, pokeapiGetByID, mysqlCreate)

	want := pokemon.MockPokemon()
	got, err := searchByID(ctx, ID)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestSearchByID_successWhenPokemonIsNotInTheDatabase(t *testing.T) {
	ID := 1
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.Pokemon{}, nil)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokeapiGetByID := pokeapi.MockGetByID(pokeapi.MockPokemon(), nil)

	ctx := context.Background()
	searchByID := pokemon.MakeSearchByID(mysqlSearchByID, pokeapiGetByID, mysqlCreate)

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
	pokeapiGetByID := pokeapi.MockGetByID(pokeapi.Pokemon{}, nil)

	ctx := context.Background()
	searchByID := pokemon.MakeSearchByID(mysqlSearchByID, pokeapiGetByID, mysqlCreate)

	want := pokemon.ErrCantSearchPokemon
	_, got := searchByID(ctx, ID)

	assert.Equal(t, got, want)
}

func TestSearchByID_failsWhenGetByIDThrowsNotFound(t *testing.T) {
	ID := 1
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.Pokemon{}, nil)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokeapiGetByID := pokeapi.MockGetByID(pokeapi.Pokemon{}, pokeapi.ErrPokemonNotFound)

	ctx := context.Background()
	searchByID := pokemon.MakeSearchByID(mysqlSearchByID, pokeapiGetByID, mysqlCreate)

	want := pokemon.ErrPokemonNotFound
	_, got := searchByID(ctx, ID)

	assert.Equal(t, got, want)
}

func TestSearchByID_failsWhenGetByIDThrowsError(t *testing.T) {
	ID := 1
	err := errors.New("error")
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.Pokemon{}, nil)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokeapiGetByID := pokeapi.MockGetByID(pokeapi.Pokemon{}, err)

	ctx := context.Background()
	searchByID := pokemon.MakeSearchByID(mysqlSearchByID, pokeapiGetByID, mysqlCreate)

	want := pokemon.ErrCantGetPokemon
	_, got := searchByID(ctx, ID)

	assert.Equal(t, got, want)
}

func TestSearchByID_failsWhenCantSavePokemon(t *testing.T) {
	ID := 1
	err := errors.New("error")
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.Pokemon{}, nil)
	mysqlCreate := pokemon.MockMySQLCreate(err)
	pokeapiGetByID := pokeapi.MockGetByID(pokeapi.MockPokemon(), nil)

	ctx := context.Background()
	searchByID := pokemon.MakeSearchByID(mysqlSearchByID, pokeapiGetByID, mysqlCreate)

	want := pokemon.ErrCantCreatePokemon
	_, got := searchByID(ctx, ID)

	assert.Equal(t, got, want)
}
