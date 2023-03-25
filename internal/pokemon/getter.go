package pokemon

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rromero96/roro-lib/cmd/rest"
)

type GetPokemon func(ctx context.Context, ID int) (Pokemon, error)

const pokeApiUrl string = "https://pokeapi.co/api/v2/pokemon/%d"

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
			return Pokemon{}, ErrNotFound
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
