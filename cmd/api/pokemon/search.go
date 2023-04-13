package pokemon

import (
	"context"

	"github.com/rromero96/PI-Pokemon/internal/pokemon"
	"github.com/rromero96/roro-lib/cmd/log"
)

type (
	// SearchTypes search the pokemon types in the db, if there are not existent, it looks for them in the pokeapi and saves them in the db
	SearchTypes func(ctx context.Context) ([]Type, error)
)

// MakeSearchTypes creates a SearchTypes function
func MakeSearchTypes(mysqlSearchTypes MySQLSearchTypes, searchPokemonTypes pokemon.SearchTypes, mysqlCreateTypes MySQLCreateType) SearchTypes {
	return func(ctx context.Context) ([]Type, error) {
		types, err := mysqlSearchTypes(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return []Type{}, ErrCantSearchTypes
		}

		if len(types) == 0 {
			pokemonTypes, err := searchPokemonTypes(ctx)
			if err != nil {
				log.Error(ctx, err.Error())
				return []Type{}, ErrCantSearchPokemonTypes
			}
			types = toTypesSlice(pokemonTypes.Types)

			if err = mysqlCreateTypes(ctx, types); err != nil {
				log.Error(ctx, err.Error())
				return []Type{}, ErrCantSaveTypes
			}

			return types, nil
		}

		return types, nil
	}
}
