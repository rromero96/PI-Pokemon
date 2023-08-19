package pokemon

import (
	"context"

	"github.com/rromero96/PI-Pokemon/internal/pokemon"
	"github.com/rromero96/roro-lib/log"
)

type (
	// SearchByID searchs a pokemon by ID
	SearchByID func(ctx context.Context, ID int) (Pokemon, error)

	// SearchTypes search the pokemon types in the db, if there are not existent, it looks for them in the pokeapi and saves them in the db
	SearchTypes func(ctx context.Context) ([]Type, error)
)

// MakeSearchByID creates a new SearchById function
func MakeSearchByID(mysqlSearchByID MySQLSearchByID, searchPokemon pokemon.Search, mysqlCreate MySQLCreate) SearchByID {
	return func(ctx context.Context, ID int) (Pokemon, error) {
		pokemon, err := mysqlSearchByID(ctx, ID)
		if err != nil {
			log.Error(ctx, err.Error())
			return Pokemon{}, ErrCantSearchPokemon
		}

		if pokemon.ID == 0 {
			poke, err := searchPokemon(ctx, &ID, nil)
			if err != nil {
				log.Error(ctx, err.Error())
				return Pokemon{}, ErrCantSearchPokemonApi
			}

			pokemon := toPokemon(poke)

			if err = mysqlCreate(ctx, pokemon); err != nil {
				log.Error(ctx, err.Error())
				return Pokemon{}, ErrCantCreatePokemon
			}

			return pokemon, nil
		}

		return pokemon, nil
	}
}

// MakeSearchTypes creates a new SearchTypes function
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
