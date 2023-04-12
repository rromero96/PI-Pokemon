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
	querySearchTypesMock string = "SELECT id, name FROM type"
)

func TestMakeMySQLSearchType_success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(querySearchTypesMock)
	id, name := 1, "electric"

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name)
	mock.ExpectQuery(querySearchTypesMock).WillReturnRows(rows)

	got, err := pokemon.MakeMySQLSearchTypes(db)

	assert.Nil(t, err)
	assert.NotNil(t, got)
}

func TestMySQLSearchType_success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare(querySearchTypesMock)
	id, name := 1, "electric"

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name)
	mock.ExpectQuery(querySearchTypesMock).WillReturnRows(rows)
	mysqlSearchType, _ := pokemon.MakeMySQLSearchTypes(db)
	ctx := context.Background()
	types := pokemon.MockTypes()
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
	id, name := 1, "electric"

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
	id, name := 1, "electric"

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
	id, name := 1, "electric"

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name).RowError(0, errors.New("some error"))
	mock.ExpectQuery(querySearchTypesMock).WillReturnRows(rows)
	mysqlSearchType, _ := pokemon.MakeMySQLSearchTypes(db)
	ctx := context.Background()

	want := pokemon.ErrCantReadRows
	_, got := mysqlSearchType(ctx)

	assert.Equal(t, want, got)
}
