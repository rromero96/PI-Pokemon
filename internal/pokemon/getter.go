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
	// SearchPokemon retrieves a Pokemon by id or name
	SearchPokemon func(context.Context, *int, *string) (Pokemon, error)

	// SearchTypes retrieves the pokemon types
	SearchTypes func(context.Context) (PokemonTypes, error)
)

// MakeGetPokemons creates a new SearchPokemon function
func MakeSearchPokemon(httpClient rusty.Requester) (SearchPokemon, error) {
	const domain string = "https://pokeapi.co"
	const path string = "/api/v2/pokemon/{key}"

	url := rusty.URL(domain, path)
	endpoint, err := rusty.NewEndpoint(httpClient, url)
	if err != nil {
		return nil, rusty.ErrCantCreateRustyEndpoint
	}

	return func(ctx context.Context, ID *int, Name *string) (Pokemon, error) {
		var requestOpts []rusty.RequestOption
		if ID != nil {
			requestOpts = append(requestOpts, rusty.WithParam("key", *ID))
		}

		if Name != nil {
			requestOpts = append(requestOpts, rusty.WithParam("key", *Name))
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
		case http.StatusNotFound:
			return Pokemon{}, ErrPokemonNotFound
		default:
			log.Error(ctx, fmt.Sprintf("Error: %s - Body: %s - StatusCode: %d", err.Error(), response.Body, response.StatusCode))
			return Pokemon{}, ErrCantGetPokemon
		}
	}, nil
}

// MakeGetTypes creates a new SearchTypes function
func MakeSearchTypes(httpClient rusty.Requester) (SearchTypes, error) {
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
