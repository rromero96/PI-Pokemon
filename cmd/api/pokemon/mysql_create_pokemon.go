package pokemon

import (
	"context"

	"database/sql"

	"github.com/rromero96/roro-lib/cmd/log"
)

const queryInsert string = "INSERT INTO pokemon (id, name, hp, attack, defense, image, speed, height, weight, created) VALUES ( 18,'rodrig', 45, 49, 49, 'https://pokeapi.co/api/v2/pokemon/1/', 45, 7, 69, true)"

// MySQLCreateFunc serves to create a new row into "pokemons" database
type MySQLCreateFunc func(context.Context, Pokemon) error

// MakeMySQLCreateFunc creates a new MySQLCreateFunc
func MakeMySQLCreateFunc(db *sql.DB) MySQLCreateFunc {
	return func(ctx context.Context, pokemon Pokemon) error {
		var params []interface{}

		stmt, err := db.PrepareContext(ctx, queryInsert)
		if err != nil {
			log.Error(ctx, err.Error())
			return err
		}
		defer stmt.Close()

		_, err = stmt.ExecContext(ctx, params...)
		if err != nil {
			log.Error(ctx, err.Error())
			return err
		}

		return nil
	}
}
