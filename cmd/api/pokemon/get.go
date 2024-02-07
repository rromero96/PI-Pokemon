package pokemon

import (
	"context"
	"errors"

	"github.com/rromero96/roro-lib/log"

	"github.com/rromero96/PI-Pokemon/internal/pokeapi"
)

type (
	// GetByID gets a pokemon by ID
	GetByID func(ctx context.Context, ID int) (Pokemon, error)

	// GetTypes gets the pokemon types in the db, if there are not existent, it looks for them in the pokeapi and saves them in the db
	GetTypes func(ctx context.Context) ([]Type, error)
)

// MakeGetByID creates a new GetByID function
func MakeGetByID(mysqlSearchByID MySQLSearchByID, getByID pokeapi.GetByID, mysqlCreate MySQLCreate) GetByID {
	return func(ctx context.Context, ID int) (Pokemon, error) {
		pokemon, err := mysqlSearchByID(ctx, ID)
		if err != nil {
			log.Error(ctx, err.Error())
			return Pokemon{}, ErrCantGetPokemon
		}

		if pokemon.ID == 0 {
			poke, err := getByID(ctx, ID)
			if err != nil {
				log.Error(ctx, err.Error())
				if errors.Is(err, pokeapi.ErrPokemonNotFound) {
					return Pokemon{}, ErrPokemonNotFound
				}

				return Pokemon{}, ErrCantGetApiPokemon
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

// MakeGetTypes creates a new GetTypes function
func MakeGetTypes(mysqlSearchTypes MySQLSearchTypes, getTypes pokeapi.GetTypes, mysqlCreateTypes MySQLCreateType) GetTypes {
	return func(ctx context.Context) ([]Type, error) {
		types, err := mysqlSearchTypes(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return []Type{}, ErrCantGetTypes
		}

		if len(types) == 0 {
			pokemonTypes, err := getTypes(ctx)
			if err != nil {
				log.Error(ctx, err.Error())
				return []Type{}, ErrCantGetApiTypes
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
