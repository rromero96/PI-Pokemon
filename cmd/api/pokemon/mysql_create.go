package pokemon

import (
	"context"
	"strings"

	"database/sql"

	"github.com/rromero96/roro-lib/log"
)

type (
	// MySQLCreate serves to create a new row into "pokemons" schema
	MySQLCreate func(ctx context.Context, pokemon Pokemon) error

	// MySQLAdd serves to add a new relationships between "pokemons" and "types" schemas
	MySQLAdd func(ctx context.Context, ID int, types []Type) error

	// MySQLCreate serves to create a new row into "type" schema
	MySQLCreateType func(ctx context.Context, types []Type) error
)

// MakeMySQLCreate creates a new MySQLCreate
func MakeMySQLCreate(db *sql.DB, addTypes MySQLAdd) MySQLCreate {
	var query string = "INSERT INTO pokemon (id, name, hp, attack, defense, image, speed, height, weight, created) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	return func(ctx context.Context, pokemon Pokemon) error {
		stmt, err := db.PrepareContext(ctx, query)
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
			log.Error(ctx, err.Error())
			return ErrCantGetLastID
		}

		if err := addTypes(ctx, int(id), pokemon.Types); err != nil {
			log.Error(ctx, err.Error())
			return ErrCantAddTypes
		}

		return nil
	}
}

// MakeMySQLAdd creates a new MySQLAdd
func MakeMySQLAdd(db *sql.DB) MySQLAdd {
	return func(ctx context.Context, ID int, types []Type) error {
		var query string = "INSERT INTO pokemon_type (pokemon_id, type_name) VALUES "
		var inserts []string
		var params []interface{}

		for _, t := range types {
			inserts = append(inserts, "(?, ?)")
			params = append(params, ID, t.Name)
		}

		queryVals := strings.Join(inserts, ",")
		query = query + queryVals

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

// MakeMySQLCreateType creates a new MySQLCreateType
func MakeMySQLCreateType(db *sql.DB) MySQLCreateType {
	return func(ctx context.Context, types []Type) error {
		var query string = "INSERT INTO type (id, name) VALUES "
		var inserts []string
		var params []interface{}

		for _, t := range types {
			inserts = append(inserts, "(?, ?)")
			params = append(params, t.ID, t.Name)
		}

		queryVals := strings.Join(inserts, ",")
		query = query + queryVals

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
