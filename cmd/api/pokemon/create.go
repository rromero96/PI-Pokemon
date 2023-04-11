package pokemon

import (
	"context"
	"strings"

	"database/sql"

	"github.com/rromero96/roro-lib/cmd/log"
)

const (
	queryCreate string = "INSERT INTO pokemon (id, name, hp, attack, defense, image, speed, height, weight, created) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	queryAdd    string = "INSERT INTO pokemon_type (pokemon_id, type_name) VALUES "
)

type (
	// MySQLCreate serves to create a new row into "pokemons" database
	MySQLCreate func(ctx context.Context, pokemon Pokemon) error

	// MySQLAdd serves to add a new relationships between "pokemons" and "types" schemas
	MySQLAdd func(ctx context.Context, ID int, types []Type) error
)

// MakeMySQLCreate creates a new MySQLCreate
func MakeMySQLCreate(db *sql.DB, addTypes MySQLAdd) MySQLCreate {
	return func(ctx context.Context, pokemon Pokemon) error {
		stmt, err := db.PrepareContext(ctx, queryCreate)
		if err != nil {
			log.Error(ctx, err.Error())
			return ErrCantPrepareStatement
		}
		defer stmt.Close()

		p, err := stmt.ExecContext(ctx, pokemon.ID, pokemon.Name, pokemon.HP, pokemon.Attack, pokemon.Defense, pokemon.Image, pokemon.Speed, pokemon.Height, pokemon.Weight, pokemon.Created)
		if err != nil {
			log.Error(ctx, err.Error())
			return ErrCantRunQuery
		}

		id, err := p.LastInsertId()
		if err != nil {
			return ErrCantGetLastID
		}

		err = addTypes(ctx, int(id), pokemon.Types)
		if err != nil {
			return ErrCantAddTypes
		}

		return nil
	}
}

// MakeMySQLAdd creates a new MySQLAdd
func MakeMySQLAdd(db *sql.DB) MySQLAdd {
	return func(ctx context.Context, ID int, types []Type) error {
		var inserts []string
		var params []interface{}

		for _, t := range types {
			inserts = append(inserts, "(?, ?)")
			params = append(params, ID, t.Name)
		}

		queryVals := strings.Join(inserts, ",")
		query := queryAdd + queryVals

		stmt, err := db.PrepareContext(ctx, query)
		if err != nil {
			log.Error(ctx, err.Error())
			return ErrCantPrepareStatement
		}
		defer stmt.Close()

		_, err = stmt.ExecContext(ctx, params...)
		if err != nil {
			log.Error(ctx, err.Error())
			return ErrCantRunQuery
		}

		return nil
	}
}
