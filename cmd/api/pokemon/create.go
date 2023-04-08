package pokemon

import (
	"context"

	"database/sql"

	"github.com/rromero96/roro-lib/cmd/log"
)

const queryInsert string = "INSERT INTO pokemon (id, name, hp, attack, defense, image, speed, height, weight, created) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

// MySQLCreateFunc serves to create a new row into "pokemons" database
type MySQLCreate func(ctx context.Context, pokemon Pokemon) error

// MakeMySQLCreate creates a new MySQLCreate
func MakeMySQLCreate(db *sql.DB) MySQLCreate {
	return func(ctx context.Context, pokemon Pokemon) error {
		stmt, err := db.PrepareContext(ctx, queryInsert)
		if err != nil {
			log.Error(ctx, err.Error())
			return ErrCantPrepareStatement
		}

		_, err = stmt.ExecContext(ctx, pokemon.ID, pokemon.Name, pokemon.HP, pokemon.Attack, pokemon.Defense, pokemon.Image, pokemon.Speed, pokemon.Height, pokemon.Weight, pokemon.Created)
		if err != nil {
			log.Error(ctx, err.Error())
			return ErrCantRunQuery
		}
		defer stmt.Close()

		return nil
	}
}
