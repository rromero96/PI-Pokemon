package pokemon_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rromero96/PI-Pokemon/cmd/api/pokemon"
	"github.com/stretchr/testify/assert"
)

const (
	querySearchTypesMock       string = "SELECT id, name FROM type ORDER BY id ASC"
	querySearchPokemonByIDMock string = "SELECT id, name, hp, attack, defense, image, speed, height, weight, custom, \\(SELECT type_name FROM pokemon_type WHERE pokemon_id = id ORDER BY type_name LIMIT 1\\) AS type_1, \\(SELECT type_name FROM pokemon_type WHERE pokemon_id = id ORDER BY type_name LIMIT 1,1\\) AS type_2 FROM pokemon WHERE id = ?"
)

func TestMakeMySQLSearchType_success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(querySearchTypesMock)
	id, name := 1, "grass"

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name)
	mock.ExpectQuery(querySearchTypesMock).WillReturnRows(rows)

	got, err := pokemon.MakeMySQLSearchTypes(db)

	assert.Nil(t, err)
	assert.NotNil(t, got)
}

func TestMySQLSearchType_success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(querySearchTypesMock)
	id, name := 1, "grass"

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name)
	mock.ExpectQuery(querySearchTypesMock).WillReturnRows(rows)
	mysqlSearchType, _ := pokemon.MakeMySQLSearchTypes(db)
	ctx := context.Background()
	types := []pokemon.Type{pokemon.MockTypes()[0]}
	types[0].ID = 1

	want := types
	got, err := mysqlSearchType(ctx)

	assert.Nil(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, want, got)
}

func TestMySQLSearchType_failsWhenCantPrepareStatement(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(pokemon.ErrCantPrepareStatement.Error())
	id, name := 1, "grass"

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name)
	mock.ExpectQuery(querySearchTypesMock).WillReturnRows(rows)
	mysqlSearchType, _ := pokemon.MakeMySQLSearchTypes(db)
	ctx := context.Background()

	want := pokemon.ErrCantPrepareStatement
	_, got := mysqlSearchType(ctx)

	assert.Equal(t, want, got)
}

func TestMySQLSearchType_failsWhenCantRunQuery(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(querySearchTypesMock)
	id, name := 1, "grass"

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name)
	mock.ExpectQuery(pokemon.ErrCantRunQuery.Error()).WillReturnRows(rows)
	mysqlSearchType, _ := pokemon.MakeMySQLSearchTypes(db)
	ctx := context.Background()

	want := pokemon.ErrCantRunQuery
	_, got := mysqlSearchType(ctx)

	assert.Equal(t, want, got)
}

func TestMySQLSearchType_failsWhenCantScanRowResult(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(querySearchTypesMock)

	rows := sqlmock.NewRows([]string{"invalid column list"}).AddRow("some value")
	mock.ExpectQuery(querySearchTypesMock).WillReturnRows(rows)
	mysqlSearchType, _ := pokemon.MakeMySQLSearchTypes(db)
	ctx := context.Background()

	want := pokemon.ErrCantScanRowResult
	_, got := mysqlSearchType(ctx)

	assert.Equal(t, want, got)
}

func TestMySQLSearchType_failsWhenRowResultHasError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(querySearchTypesMock)
	id, name := 1, "grass"

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name).RowError(0, errors.New("some error"))
	mock.ExpectQuery(querySearchTypesMock).WillReturnRows(rows)
	mysqlSearchType, _ := pokemon.MakeMySQLSearchTypes(db)
	ctx := context.Background()

	want := pokemon.ErrCantReadRows
	_, got := mysqlSearchType(ctx)

	assert.Equal(t, want, got)
}

func TestMakeMySQLSearchByID_success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(querySearchPokemonByIDMock)
	id, name, hp, attack, defense, image, speed, height, weight, custom, type_1, type_2 := 1, "bulbasaur", 100, 100, 100, "image", 100, 100, 100, false, "grass", "poison"

	rows := sqlmock.NewRows([]string{"id", "name", "hp", "attack", "defense", "image", "speed", "height", "weight", "custom", "type_1", "type_2"}).AddRow(id, name, hp, attack, defense, image, speed, height, weight, custom, type_1, type_2)
	mock.ExpectQuery(querySearchPokemonByIDMock).WillReturnRows(rows)

	got, err := pokemon.MakeMySQLSearchByID(db)

	assert.Nil(t, err)
	assert.NotNil(t, got)
}

func TestMySQLSearchByID_success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(querySearchPokemonByIDMock)
	id, name, hp, attack, defense, image, speed, height, weight, custom, type_1, type_2 := 1, "bulbasaur", 100, 100, 100, "image", 100, 100, 100, false, "grass", "poison"

	rows := sqlmock.NewRows([]string{"id", "name", "hp", "attack", "defense", "image", "speed", "height", "weight", "custom", "type_1", "type_2"}).AddRow(id, name, hp, attack, defense, image, speed, height, weight, custom, type_1, type_2)
	mock.ExpectQuery(querySearchPokemonByIDMock).WillReturnRows(rows)
	mysqlSearchByID, _ := pokemon.MakeMySQLSearchByID(db)
	ctx := context.Background()

	want := pokemon.MockPokemon()
	got, err := mysqlSearchByID(ctx, id)

	assert.Nil(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, want, got)
}

func TestMySQLSearchByID_failsWhenCantPrepareStatement(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(pokemon.ErrCantPrepareStatement.Error())
	id, name, hp, attack, defense, image, speed, height, weight, custom, type_1, type_2 := 1, "bulbasaur", 100, 100, 100, "image", 100, 100, 100, false, "grass", "poison"

	rows := sqlmock.NewRows([]string{"id", "name", "hp", "attack", "defense", "image", "speed", "height", "weight", "custom", "type_1", "type_2"}).AddRow(id, name, hp, attack, defense, image, speed, height, weight, custom, type_1, type_2)
	mock.ExpectQuery(querySearchPokemonByIDMock).WillReturnRows(rows)
	mysqlSearchByID, _ := pokemon.MakeMySQLSearchByID(db)
	ctx := context.Background()

	want := pokemon.ErrCantPrepareStatement
	_, got := mysqlSearchByID(ctx, id)

	assert.Equal(t, want, got)
}

func TestMySQLSearchByID_failsWhenCantRunQuery(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(querySearchPokemonByIDMock)
	id, name, hp, attack, defense, image, speed, height, weight, custom, type_1, type_2 := 1, "bulbasaur", 100, 100, 100, "image", 100, 100, 100, false, "grass", "poison"

	rows := sqlmock.NewRows([]string{"id", "name", "hp", "attack", "defense", "image", "speed", "height", "weight", "custom", "type_1", "type_2"}).AddRow(id, name, hp, attack, defense, image, speed, height, weight, custom, type_1, type_2)
	mock.ExpectQuery(pokemon.ErrCantRunQuery.Error()).WillReturnRows(rows)
	mysqlSearchByID, _ := pokemon.MakeMySQLSearchByID(db)
	ctx := context.Background()

	want := pokemon.ErrCantRunQuery
	_, got := mysqlSearchByID(ctx, id)

	assert.Equal(t, want, got)
}

func TestMySQLSearchByID_failsWhenCantScanRowResult(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(querySearchPokemonByIDMock)
	id := 1

	rows := sqlmock.NewRows([]string{"invalid column list"}).AddRow("some value")
	mock.ExpectQuery(querySearchPokemonByIDMock).WillReturnRows(rows)
	mysqlSearchByID, _ := pokemon.MakeMySQLSearchByID(db)
	ctx := context.Background()

	want := pokemon.ErrCantScanRowResult
	_, got := mysqlSearchByID(ctx, id)

	assert.Equal(t, want, got)
}

func TestMySQLSearchByID_failsWhenRowResultHasError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(querySearchPokemonByIDMock)
	id, name, hp, attack, defense, image, speed, height, weight, custom, type_1, type_2 := 1, "bulbasaur", 100, 100, 100, "image", 100, 100, 100, false, "grass", "poison"

	rows := sqlmock.NewRows([]string{"id", "name", "hp", "attack", "defense", "image", "speed", "height", "weight", "custom", "type_1", "type_2"}).AddRow(id, name, hp, attack, defense, image, speed, height, weight, custom, type_1, type_2).RowError(0, errors.New("some error"))
	mock.ExpectQuery(querySearchPokemonByIDMock).WillReturnRows(rows)
	mysqlSearchByID, _ := pokemon.MakeMySQLSearchByID(db)
	ctx := context.Background()

	want := pokemon.ErrCantReadRows
	_, got := mysqlSearchByID(ctx, id)

	assert.Equal(t, want, got)
}
