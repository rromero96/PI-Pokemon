package pokemon_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/rromero96/PI-Pokemon/cmd/api/pokemon"
)

const (
	queryCreateMock = "INSERT INTO pokemon \\(id, name, hp, attack, defense, image, speed, height, weight, created\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)"
	queryAddMock    = "INSERT INTO pokemon_type \\(pokemon_id, type_name\\) VALUES \\(\\?, \\?\\)"
)

func TestMakeMySQLCreate_success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(queryCreateMock)
	mock.ExpectExec(queryCreateMock).WillReturnResult(sqlmock.NewResult(1, 2))

	got := pokemon.MakeMySQLCreate(db)

	assert.NotNil(t, got)
}

func TestMySQLCreate_success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(queryCreateMock)
	mock.ExpectExec(queryCreateMock).WillReturnResult(sqlmock.NewResult(1, 2))

	mysqlCreate := pokemon.MakeMySQLCreate(db)
	ctx := context.Background()

	got := mysqlCreate(ctx, pokemon.MockPokemon())

	assert.Nil(t, got)
}

func TestMySQLCreate_failsWhenCantPrepareStatement(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare("invalid statement")
	mock.ExpectExec(queryCreateMock).WillReturnResult(sqlmock.NewResult(1, 2))

	mysqlCreate := pokemon.MakeMySQLCreate(db)
	ctx := context.Background()

	want := pokemon.ErrCantPrepareStatement
	got := mysqlCreate(ctx, pokemon.MockPokemon())

	assert.Equal(t, want, got)
}

func TestMySQLCreate_failsWhenCantRunQuery(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(queryCreateMock)
	mock.ExpectExec(queryCreateMock).WillReturnError(errors.New("some error"))

	mysqlCreate := pokemon.MakeMySQLCreate(db)
	ctx := context.Background()

	want := pokemon.ErrCantRunQuery
	got := mysqlCreate(ctx, pokemon.MockPokemon())

	assert.Equal(t, want, got)
}

func TestMakeMySQLAdd_success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(queryAddMock)
	mock.ExpectExec(queryAddMock).WillReturnResult(sqlmock.NewResult(1, 2))

	got := pokemon.MakeMySQLAdd(db)

	assert.NotNil(t, got)
}

func TestMySQLAdd_success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(queryAddMock)
	mock.ExpectExec(queryAddMock).WillReturnResult(sqlmock.NewResult(1, 2))

	mysqlAdd := pokemon.MakeMySQLAdd(db)
	ctx := context.Background()

	got := mysqlAdd(ctx, pokemon.MockPokemon().ID, pokemon.MockTypes())

	assert.Nil(t, got)
}

func TestMySQLAdd_failsWhenCantPrepareStatement(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare("invalid statement")
	mock.ExpectExec(queryAddMock).WillReturnResult(sqlmock.NewResult(1, 2))

	mysqlAdd := pokemon.MakeMySQLAdd(db)
	ctx := context.Background()
	want := pokemon.ErrCantPrepareStatement
	got := mysqlAdd(ctx, pokemon.MockPokemon().ID, pokemon.MockTypes())

	assert.Equal(t, want, got)
}

func TestMySQLAdd_failsWhenCantRunQuery(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(queryAddMock)
	mock.ExpectExec(queryAddMock).WillReturnError(errors.New("some error"))

	mysqlAdd := pokemon.MakeMySQLAdd(db)
	ctx := context.Background()

	want := pokemon.ErrCantRunQuery
	got := mysqlAdd(ctx, pokemon.MockPokemon().ID, pokemon.MockTypes())

	assert.Equal(t, want, got)
}
