package pokemon_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/rromero96/PI-Pokemon/cmd/api/pokemon"
)

const queryInsertMock = "INSERT INTO pokemon \\(id, name, hp, attack, defense, image, speed, height, weight, created\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)"

func TestMakeMySQLCreate_success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(queryInsertMock)
	mock.ExpectExec(queryInsertMock).WillReturnResult(sqlmock.NewResult(1, 2))

	got := pokemon.MakeMySQLCreate(db)

	assert.NotNil(t, got)
}

func TestMySQLCreate_success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(queryInsertMock)
	mock.ExpectExec(queryInsertMock).WillReturnResult(sqlmock.NewResult(1, 2))

	mysqlCreate := pokemon.MakeMySQLCreate(db)
	ctx := context.Background()

	got := mysqlCreate(ctx, pokemon.MockPokemon())

	assert.Nil(t, got)
}

func TestMySQLCreate_failsWhenCantPrepareStatement(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare("invalid statement")
	mock.ExpectExec(queryInsertMock).WillReturnResult(sqlmock.NewResult(1, 2))

	mysqlCreate := pokemon.MakeMySQLCreate(db)
	ctx := context.Background()

	want := pokemon.ErrCantPrepareStatement
	got := mysqlCreate(ctx, pokemon.MockPokemon())

	assert.Equal(t, want, got)
}

func TestMySQLCreate_failsWhenCantRunQuery(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(queryInsertMock)
	mock.ExpectExec(queryInsertMock).WillReturnError(errors.New("some error"))

	mysqlCreate := pokemon.MakeMySQLCreate(db)
	ctx := context.Background()

	want := pokemon.ErrCantRunQuery
	got := mysqlCreate(ctx, pokemon.MockPokemon())

	assert.Equal(t, want, got)
}
