package pokemon

import (
	"context"
	"database/sql"

	"github.com/rromero96/roro-lib/cmd/log"
)

const (
	querySearchTypes string = "SELECT id, name FROM type"
)

type (
	// MYSQLSearchTypes performs a SELECT into pokemons database to seek the pokemon types
	MySQLSearchTypes func(ctx context.Context) ([]Type, error)
)

func MakeMySQLSearchTypes(db *sql.DB) (MySQLSearchTypes, error) {
	return func(ctx context.Context) ([]Type, error) {
		stmt, err := db.PrepareContext(ctx, querySearchTypes)
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
		if rows.Err() != nil {
			log.Error(ctx, err.Error())
			return nil, ErrCantReadRows
		}

		return types, nil
	}, nil
}
