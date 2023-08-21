package pokemon

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rromero96/roro-lib/log"
	"github.com/rromero96/roro-lib/rusty"
)

type (
	// GetByID retrieves a Pokemon by id
	GetByID func(context.Context, int) (Pokemon, error)

	// GetTypes retrieves the pokemon types
	GetTypes func(context.Context) (PokemonTypes, error)
)

// MakeGetByID creates a new GetByIDfunction
func MakeGetByID(httpClient rusty.Requester) (GetByID, error) {
	const domain string = "https://pokeapi.co"
	const path string = "/api/v2/pokemon/{pokemon_id}"

	url := rusty.URL(domain, path)
	endpoint, err := rusty.NewEndpoint(httpClient, url)
	if err != nil {
		return nil, rusty.ErrCantCreateRustyEndpoint
	}

	return func(ctx context.Context, ID int) (Pokemon, error) {
		requestOpts := []rusty.RequestOption{
			rusty.WithParam("pokemon_id", ID),
		}
		response, err := endpoint.Post(ctx, requestOpts...)
		if response == nil && err != nil {
			log.Error(ctx, ErrCantPerformGet.Error(), log.String("response error:", err.Error()))
			return Pokemon{}, ErrCantPerformGet
		}

		switch response.StatusCode {
		case http.StatusOK:
			var pokemon Pokemon
			if json.Unmarshal(response.Body, &pokemon) != nil {
				return Pokemon{}, ErrUnmarshalResponse
			}
			return pokemon, nil
		case http.StatusBadRequest, http.StatusNotFound:
			return Pokemon{}, ErrPokemonNotFound
		default:
			log.Error(ctx, fmt.Sprintf("Error: %s - Body: %s - StatusCode: %d", err.Error(), response.Body, response.StatusCode))
			return Pokemon{}, ErrCantGetPokemon
		}
	}, nil
}

// MakeGetTypes creates a new GetTypes function
func MakeGetTypes(httpClient rusty.Requester) (GetTypes, error) {
	const domain string = "https://pokeapi.co"
	const path string = "/api/v2/type"

	url := rusty.URL(domain, path)
	endpoint, err := rusty.NewEndpoint(httpClient, url)
	if err != nil {
		return nil, rusty.ErrCantCreateRustyEndpoint
	}

	return func(ctx context.Context) (PokemonTypes, error) {
		response, err := endpoint.Post(ctx)
		if response == nil && err != nil {
			log.Error(ctx, ErrCantPerformGet.Error(), log.String("response error:", err.Error()))
			return PokemonTypes{}, ErrCantPerformGet
		}

		switch response.StatusCode {
		case http.StatusOK:
			var types PokemonTypes
			if json.Unmarshal(response.Body, &types) != nil {
				return PokemonTypes{}, ErrUnmarshalResponse
			}
			return types, nil
		case http.StatusNotFound:
			return PokemonTypes{}, ErrTypesNotFound
		default:
			log.Error(ctx, fmt.Sprintf("Error: %s - Body: %s - StatusCode: %d", err.Error(), response.Body, response.StatusCode))
			return PokemonTypes{}, ErrCantGetTypes
		}
	}, nil
}
