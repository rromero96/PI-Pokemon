package pokemon

import (
	"errors"
	"net/http"

	"github.com/rromero96/roro-lib/web"
)

const ParamPokemonID string = "pokemon_id"

// GetV1 performs a get to obtain all the pokemons
func GetV1() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

// GetByIDV1 performs a get to obtain a pokemon by ID
func GetByIDV1(getByID GetByID) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := web.Params(r).Int(ParamPokemonID)
		if err != nil {
			return web.NewError(http.StatusBadRequest, InvalidID)
		}

		pokemon, err := getByID(r.Context(), id)
		if err != nil {
			switch {
			case errors.Is(err, ErrPokemonNotFound):
				return web.NewError(http.StatusNotFound, NotFound)
			default:
				return web.NewError(http.StatusInternalServerError, CantSearchPokemon)
			}
		}

		return web.EncodeJSON(w, pokemon.toDTO(), http.StatusOK)
	}
}

// CreateV1 perfoms a pokemon creation
func CreateV1(createPokemon Create) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var body PokemonDTO
		if web.DecodeJSON(r, &body) != nil || body.validate() != nil {
			return web.NewError(http.StatusBadRequest, InvalidBody)
		}

		err := createPokemon(r.Context(), body.toDomain())
		if err != nil {
			switch {
			case errors.Is(err, ErrInvalidPokemon):
				return web.NewError(http.StatusBadRequest, InvalidPokemon)
			default:
				return web.NewError(http.StatusInternalServerError, CantCreatePokemon)
			}
		}

		return web.EncodeJSON(w, "", http.StatusNoContent)
	}
}

// GetTypesV1 performs a get to obtain all pokemon types
func GetTypesV1(getTypes GetTypes) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		types, err := getTypes(r.Context())
		if err != nil {
			return web.NewError(http.StatusInternalServerError, CantGetTypes)
		}

		return web.EncodeJSON(w, toTypesDTO(types), http.StatusOK)
	}
}
