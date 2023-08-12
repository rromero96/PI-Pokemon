package pokemon

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rromero96/roro-lib/cmd/rest"
)

type (
	// SearchPokemon retrieves a Pokemon by id.
	SearchPokemon func(context.Context, int) (Pokemon, error)

	// SearchTypes retrieves the pokemon types
	SearchTypes func(context.Context) (PokemonTypes, error)
)

const (
	pokeApiUrl string = "/api/v2/pokemon/%d"
	typesUrl   string = "/api/v2/type"
)

// MakeGetPokemons creates a new SearchPokemon function
func MakeSearchPokemon(restGetFunc rest.GetFunc) (SearchPokemon, error) {
	return func(ctx context.Context, ID int) (Pokemon, error) {
		url := fmt.Sprintf(pokeApiUrl, ID)
		response := restGetFunc(ctx, url)

		switch response.StatusCode() {
		case http.StatusOK:
			var pokemon Pokemon
			if json.Unmarshal(response.Bytes(), &pokemon) != nil {
				return Pokemon{}, ErrUnmarshalResponse
			}
			return pokemon, nil
		case http.StatusNotFound:
			return Pokemon{}, ErrPokemonNotFound
		default:
			return Pokemon{}, rest.RequestError{
				Method:          http.MethodGet,
				URL:             url,
				StatusCode:      response.StatusCode(),
				ResponsePayload: response.String(),
			}
		}
	}, nil
}

// MakeGetTypes creates a new SearchTypes function
func MakeSearchTypes(restGetFunc rest.GetFunc) (SearchTypes, error) {
	return func(ctx context.Context) (PokemonTypes, error) {
		response := restGetFunc(ctx, typesUrl)

		switch response.StatusCode() {
		case http.StatusOK:
			var types PokemonTypes
			if json.Unmarshal(response.Bytes(), &types) != nil {
				return PokemonTypes{}, ErrUnmarshalResponse
			}
			return types, nil
		case http.StatusNotFound:
			return PokemonTypes{}, ErrTypesNotFound
		default:
			return PokemonTypes{}, rest.RequestError{
				Method:          http.MethodGet,
				URL:             typesUrl,
				StatusCode:      response.StatusCode(),
				ResponsePayload: response.String(),
			}
		}
	}, nil
}
