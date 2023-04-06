package pokemon

import (
	"context"

	"database/sql"

	"github.com/rromero96/roro-lib/cmd/log"
)

// ( 18,'rodrig', 45, 49, 49, 'https://pokeapi.co/api/v2/pokemon/1/', 45, 7, 69, true)
const queryInsert string = "INSERT INTO pokemon (id, name, hp, attack, defense, image, speed, height, weight, created) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"

// MySQLCreateFunc serves to create a new row into "pokemons" database
type MySQLCreate func(ctx context.Context, pokemon Pokemon) error

// MakeMySQLCreate creates a new MySQLCreate
func MakeMySQLCreate(db *sql.DB) MySQLCreate {
	return func(ctx context.Context, pokemon Pokemon) error {
		var params []interface{}

		//params = append(params, pokemon.ID, pokemon.Name, pokemon.HP, pokemon.Attack, pokemon.Defense, "https://pokeapi.co/api/v2/pokemon/1/", pokemon.Speed, pokemon.Height, pokemon.Weight, pokemon.Created)
		stmt, err := db.PrepareContext(ctx, queryInsert)
		if err != nil {
			log.Error(ctx, err.Error())
			return err
		}
		query, err := stmt.Exec(36, "aassd", 45, 49, 49, "https://pokeapi.co/api/v2/pokemon/1/", 45, 7, 69, true)
		if err != nil {
			return nil
		}
		defer stmt.Close()

		params = append(params, query)

		_, err = stmt.ExecContext(ctx, params)
		if err != nil {
			log.Error(ctx, err.Error())
			return ErrCantRunQuery
		}

		return nil
	}
}
