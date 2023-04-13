package pokemon

import (
	"errors"
	"net/http"

	"github.com/rromero96/roro-lib/cmd/web"
)

// SearchV1 performs a search to obtain all the pokemons
func SearchV1() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

// SearchVByIDV1 performs a search to obtain a pokemon by ID or by name
func SearchByIDV1() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
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
