package pokemon

import (
	"context"
)

// MockSearchPokemons mock
func MockSearchPokemons(response Pokemon, err error) SearchPokemon {
	return func(context.Context, *int, *string) (Pokemon, error) {
		return response, err
	}
}

// MockSearchTypes mock
func MockSearchTypes(response PokemonTypes, err error) SearchTypes {
	return func(context.Context) (PokemonTypes, error) {
		return response, err
	}
}

// MockPokemonAsJson mock
func MockPokemonAsJson() string {
	return `
	{
		"name": "bulbasaur",
		"id": 1,
		"height": 100,
		"weight": 100,
		"stats": [
        {
            "base_stat": 100,
            "stat": {
                "name": "hp"
            }
        }
        ],
		"types": [
			{
				"type": {
					"name": "grass"
				}
			},
			{
				"type": {
					"name": "poison"
				}
			}
		],
		"sprites": {
            "other": {
                "dream_world": {
                "front_default": "image"
                 }
            }
        }
	}
	`
}

// MockPokemon mock
func MockPokemon() Pokemon {
	return Pokemon{
		ID:     1,
		Name:   "bulbasaur",
		Height: 100,
		Weight: 100,
		Sprites: Sprites{
			Other: Other{
				DreamWorld: DreamWorld{
					FrontDefault: "image",
				},
			},
		},
		Stats: []Stats{{
			BaseStat: 100,
			Stat: Stat{
				Name: "hp",
			},
		}},
		Types: []Types{{
			Type: Type{Name: "grass"},
		}, {
			Type: Type{Name: "poison"},
		}},
	}
}

// MockTypesAsJson mock
func MockTypesAsJson() string {
	return `
	{
		"results": [
             {
                 "name": "grass"
             },
			 {
				"name": "poison"
			 }
         ]
	}
	`
}

// MockTypes mock
func MockTypes() PokemonTypes {
	return PokemonTypes{
		Types: []Type{{
			Name: "grass",
		}, {
			Name: "poison",
		}},
	}
}
