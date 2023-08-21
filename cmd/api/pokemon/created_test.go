package pokemon_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rromero96/PI-Pokemon/cmd/api/pokemon"
)

func TestMakeCreate_success(t *testing.T) {
	mysqlCreate := pokemon.MockMySQLCreate(nil)

	got := pokemon.MakeCreate(mysqlCreate)

	assert.NotNil(t, got)
}

func TestCreate_success(t *testing.T) {
	mysqlCreate := pokemon.MockMySQLCreate(nil)

	ctx := context.Background()
	create := pokemon.MakeCreate(mysqlCreate)

	got := create(ctx, pokemon.MockPokemon())

	assert.Nil(t, got)
}

func TestCreate_failsWhenMySQLCreateThrowsError(t *testing.T) {
	mysqlCreate := pokemon.MockMySQLCreate(pokemon.ErrCantPrepareStatement)

	ctx := context.Background()
	create := pokemon.MakeCreate(mysqlCreate)

	want := pokemon.ErrCantCreatePokemon
	got := create(ctx, pokemon.MockPokemon())

	assert.Equal(t, got, want)
}

func TestCreate_failsWhenMySQLCreateThrowsErrorOnQuery(t *testing.T) {
	mysqlCreate := pokemon.MockMySQLCreate(pokemon.ErrCantRunQuery)

	ctx := context.Background()
	create := pokemon.MakeCreate(mysqlCreate)

	want := pokemon.ErrInvalidPokemon
	got := create(ctx, pokemon.MockPokemon())

	assert.Equal(t, got, want)
}
