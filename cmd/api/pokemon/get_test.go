package pokemon_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rromero96/PI-Pokemon/cmd/api/pokemon"
	"github.com/rromero96/PI-Pokemon/internal/pokeapi"
)

func TestGetTypes_success(t *testing.T) {
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(pokemon.MockTypes(), nil)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(nil)
	pokeapiGetTypes := pokeapi.MockGetTypes(pokeapi.MockPokemonTypesDTO(), nil)

	ctx := context.Background()
	getTypes := pokemon.MakeGetTypes(mysqlSearchTypes, pokeapiGetTypes, mysqlCreateTypes)

	want := pokemon.MockTypes()
	got, err := getTypes(ctx)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestGetTypes_failsWhenCantSearchTypes(t *testing.T) {
	err := errors.New("error")
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(nil, err)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(nil)
	pokeapiGetTypes := pokeapi.MockGetTypes(pokeapi.MockPokemonTypesDTO(), nil)

	ctx := context.Background()
	getTypes := pokemon.MakeGetTypes(mysqlSearchTypes, pokeapiGetTypes, mysqlCreateTypes)

	want := pokemon.ErrCantGetTypes
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}

func TestGetTypes_failsWhenCantGetPokemonTypes(t *testing.T) {
	err := errors.New("error")
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(nil, nil)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(nil)
	pokeapiGetTypes := pokeapi.MockGetTypes(pokeapi.PokemonTypesDTO{}, err)

	ctx := context.Background()
	getTypes := pokemon.MakeGetTypes(mysqlSearchTypes, pokeapiGetTypes, mysqlCreateTypes)

	want := pokemon.ErrCantGetApiTypes
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}

func TestGetTypes_failsWhenCantSaveTypes(t *testing.T) {
	err := errors.New("error")
	mysqlSearchTypes := pokemon.MockMySQLSearchTypes(nil, nil)
	mysqlCreateTypes := pokemon.MockMySQLCreateType(err)
	pokeapiGetTypes := pokeapi.MockGetTypes(pokeapi.MockPokemonTypesDTO(), nil)

	ctx := context.Background()
	getTypes := pokemon.MakeGetTypes(mysqlSearchTypes, pokeapiGetTypes, mysqlCreateTypes)

	want := pokemon.ErrCantSaveTypes
	_, got := getTypes(ctx)

	assert.Equal(t, got, want)
}

func TestGetByID_success(t *testing.T) {
	ID := 1
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.MockPokemon(), nil)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokeapiGetByID := pokeapi.MockGetByID(pokeapi.PokemonDTO{}, nil)

	ctx := context.Background()
	getByID := pokemon.MakeGetByID(mysqlSearchByID, pokeapiGetByID, mysqlCreate)

	want := pokemon.MockPokemon()
	got, err := getByID(ctx, ID)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestGetByID_successWhenPokemonIsNotInTheDatabase(t *testing.T) {
	ID := 1
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.Pokemon{}, nil)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokeapiGetByID := pokeapi.MockGetByID(pokeapi.MockPokemonDTO(), nil)

	ctx := context.Background()
	getByID := pokemon.MakeGetByID(mysqlSearchByID, pokeapiGetByID, mysqlCreate)

	want := pokemon.MockPokemon()
	got, err := getByID(ctx, ID)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestGetByID_failsWhenCantSearchByID(t *testing.T) {
	ID := 1
	err := errors.New("error")
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.Pokemon{}, err)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokeapiGetByID := pokeapi.MockGetByID(pokeapi.PokemonDTO{}, nil)

	ctx := context.Background()
	getByID := pokemon.MakeGetByID(mysqlSearchByID, pokeapiGetByID, mysqlCreate)

	want := pokemon.ErrCantGetPokemon
	_, got := getByID(ctx, ID)

	assert.Equal(t, got, want)
}

func TestGetByID_failsWhenGetByIDThrowsNotFound(t *testing.T) {
	ID := 1
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.Pokemon{}, nil)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokeapiGetByID := pokeapi.MockGetByID(pokeapi.PokemonDTO{}, pokeapi.ErrPokemonNotFound)

	ctx := context.Background()
	getByID := pokemon.MakeGetByID(mysqlSearchByID, pokeapiGetByID, mysqlCreate)

	want := pokemon.ErrPokemonNotFound
	_, got := getByID(ctx, ID)

	assert.Equal(t, got, want)
}

func TestGetByID_failsWhenGetByIDThrowsError(t *testing.T) {
	ID := 1
	err := errors.New("error")
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.Pokemon{}, nil)
	mysqlCreate := pokemon.MockMySQLCreate(nil)
	pokeapiGetByID := pokeapi.MockGetByID(pokeapi.PokemonDTO{}, err)

	ctx := context.Background()
	getByID := pokemon.MakeGetByID(mysqlSearchByID, pokeapiGetByID, mysqlCreate)

	want := pokemon.ErrCantGetPokemon
	_, got := getByID(ctx, ID)

	assert.Equal(t, got, want)
}

func TestGetByID_failsWhenCantSavePokemon(t *testing.T) {
	ID := 1
	err := errors.New("error")
	mysqlSearchByID := pokemon.MockMySQLSearchByID(pokemon.Pokemon{}, nil)
	mysqlCreate := pokemon.MockMySQLCreate(err)
	pokeapiGetByID := pokeapi.MockGetByID(pokeapi.MockPokemonDTO(), nil)

	ctx := context.Background()
	getByID := pokemon.MakeGetByID(mysqlSearchByID, pokeapiGetByID, mysqlCreate)

	want := pokemon.ErrCantCreatePokemon
	_, got := getByID(ctx, ID)

	assert.Equal(t, got, want)
}
