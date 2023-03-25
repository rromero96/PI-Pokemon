package pokemon

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rromero96/roro-lib/cmd/rest"
)

type (
	// GetPokemon retrieves a Pokemon by id.
	GetPokemon func(context.Context, int) (Pokemon, error)

	// GetTypes retrieves the pokemon types
	GetTypes func(context.Context) (PokemonTypes, error)
)

const (
	pokeApiUrl string = "https://pokeapi.co/api/v2/pokemon/%d"
	typesUrl   string = "https://pokeapi.co/api/v2/type"
)

// MakeGetPokemons creates a new GetPokemon function
func MakeGetPokemon(restGetFunc rest.GetFunc) (GetPokemon, error) {
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

// MakeGetTypes creates a new GetTypes function
func MakeGetTypes(restGetFunc rest.GetFunc) (GetTypes, error) {
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
