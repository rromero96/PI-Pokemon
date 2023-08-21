package pokemon

import (
	"context"
	"errors"

	"github.com/rromero96/roro-lib/log"
)

type (
	// Create creates a new pokemon in the db
	Create func(ctx context.Context, pokemon Pokemon) error
)

// MakeCreate creates a new Create function
func MakeCreate(mysqlCreate MySQLCreate) Create {
	return func(ctx context.Context, pokemon Pokemon) error {
		pokemon.Custom = true

		err := mysqlCreate(ctx, pokemon)
		if err != nil {
			log.Error(ctx, err.Error())
			if errors.Is(err, ErrCantRunQuery) {
				return ErrInvalidPokemon
			}

			return ErrCantCreatePokemon
		}

		return nil
	}
}
