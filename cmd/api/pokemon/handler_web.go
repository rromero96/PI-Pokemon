package pokemon

import (
	"errors"
	"net/http"

	"github.com/rromero96/roro-lib/web"
)

const ParamPokemonID string = "pokemon_id"

// SearchV1 performs a search to obtain all the pokemons
func SearchV1() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

// SearchVByIDV1 performs a search to obtain a pokemon by ID
func SearchByIDV1(searchByID SearchByID) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := web.Params(r).Int(ParamPokemonID)
		if err != nil {
			return web.NewError(http.StatusBadRequest, InvalidID)
		}

		pokemon, err := searchByID(r.Context(), id)
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
func CreateV1(createPokemon MySQLCreate) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var body PokemonDTO
		if web.DecodeJSON(r, &body) != nil || body.validate() != nil {
			return web.NewError(http.StatusBadRequest, InvalidBody)
		}

		err := createPokemon(r.Context(), body.toDomain())
		if err != nil {
			switch {
			case errors.Is(err, ErrCantRunQuery):
				return web.NewError(http.StatusBadRequest, InvalidPokemon)
			default:
				return web.NewError(http.StatusInternalServerError, CantCreatePokemon)
			}
		}

		return web.EncodeJSON(w, "", http.StatusNoContent)
	}
}

// SearchTypesV1 performs a search to obtain all pokemon types
func SearchTypesV1(searchTypes SearchTypes) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		types, err := searchTypes(r.Context())
		if err != nil {
			return web.NewError(http.StatusInternalServerError, CantGetTypes)
		}

		return web.EncodeJSON(w, toTypesDTO(types), http.StatusOK)
	}
}
