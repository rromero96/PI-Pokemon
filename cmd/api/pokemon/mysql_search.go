package pokemon

import (
	"context"
	"database/sql"

	"github.com/rromero96/roro-lib/log"
)

type (
	// MYSQLSearchByID performs a SELECT into pokemons database to seek a pokemon by ID
	MySQLSearchByID func(ctx context.Context, ID int) (Pokemon, error)

	// MYSQLSearchTypes performs a SELECT into pokemons database to seek the pokemon types
	MySQLSearchTypes func(ctx context.Context) ([]Type, error)
)

// MakeMySQLSearchByID creates a new MySQLSearchByID function
func MakeMySQLSearchByID(db *sql.DB) MySQLSearchByID {
	return func(ctx context.Context, ID int) (Pokemon, error) {
		const query string = `SELECT id, name, hp, attack, defense, image, speed, height, weight, custom, 
		(SELECT type_name FROM pokemon_type WHERE pokemon_id = id ORDER BY type_name LIMIT 1) AS type_1,
		(SELECT type_name FROM pokemon_type WHERE pokemon_id = id ORDER BY type_name LIMIT 1,1) AS type_2
		FROM pokemon
		WHERE id = ?`

		stmt, err := db.PrepareContext(ctx, query)
		if err != nil {
			log.Error(ctx, err.Error())
			return Pokemon{}, ErrCantPrepareStatement
		}
		defer stmt.Close()

		rows, err := stmt.QueryContext(ctx, ID)
		if err != nil {
			log.Error(ctx, err.Error())
			return Pokemon{}, ErrCantRunQuery
		}
		defer rows.Close()

		var pokemon Pokemon
		for rows.Next() {
			var type1, type2 sql.NullString
			if err := rows.Scan(&pokemon.ID, &pokemon.Name, &pokemon.HP, &pokemon.Attack, &pokemon.Defense, &pokemon.Image, &pokemon.Speed, &pokemon.Height, &pokemon.Weight, &pokemon.Custom, &type1, &type2); err != nil {
				log.Error(ctx, err.Error())
				return Pokemon{}, ErrCantScanRowResult
			}

			if type1.Valid {
				pokemon.Types = append(pokemon.Types, Type{Name: type1.String})
			}

			if type2.Valid {
				pokemon.Types = append(pokemon.Types, Type{Name: type2.String})
			}
		}
		if err := rows.Err(); err != nil {
			log.Error(ctx, err.Error())
			return Pokemon{}, ErrCantReadRows
		}

		return pokemon, nil
	}
}

// MakeMySQLSearchTypes creates a new MySQLSearchTypes function
func MakeMySQLSearchTypes(db *sql.DB) MySQLSearchTypes {
	return func(ctx context.Context) ([]Type, error) {
		const query string = "SELECT id, name FROM type ORDER BY id ASC"

		stmt, err := db.PrepareContext(ctx, query)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, ErrCantPrepareStatement
		}
		defer stmt.Close()

		rows, err := stmt.QueryContext(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, ErrCantRunQuery
		}
		defer rows.Close()

		var types []Type
		for rows.Next() {
			var t Type
			if err := rows.Scan(&t.ID, &t.Name); err != nil {
				log.Error(ctx, err.Error())
				return nil, ErrCantScanRowResult
			}
			types = append(types, t)
		}
		if err := rows.Err(); err != nil {
			log.Error(ctx, err.Error())
			return nil, ErrCantReadRows
		}

		return types, nil
	}
}
