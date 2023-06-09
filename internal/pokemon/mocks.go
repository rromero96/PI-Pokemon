package pokemon

import (
	"context"
)

// MockSearchPokemons mock
func MockSearchPokemons(response Pokemon, err error) SearchPokemon {
	return func(context.Context, int) (Pokemon, error) {
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
		"name": "pikachu",
		"id": 25,
		"height": 4,
		"weight": 60,
		"stats": [
        {
            "base_stat": 35,
            "stat": {
                "name": "hp"
            }
        }
        ],
		"types": [
			{
				"type": {
					"name": "electric"
				}
			}
		],
		"sprites": {
            "other": {
                "dream_world": {
                "front_default": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/dream-world/25.svg"
                 }
            }
        }
	}
	`
}

// MockPokemon mock
func MockPokemon() Pokemon {
	return Pokemon{
		ID:     25,
		Name:   "pikachu",
		Height: 4,
		Weight: 60,
		Sprites: Sprites{
			Other: Other{
				DreamWorld: DreamWorld{
					FrontDefault: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/dream-world/25.svg",
				},
			},
		},
		Stats: []Stats{{
			BaseStat: 35,
			Stat: Stat{
				Name: "hp",
			},
		}},
		Types: []Types{{
			Type: Type{
				Name: "electric",
			},
		}},
	}
}

// MockTypesAsJson mock
func MockTypesAsJson() string {
	return `
	{
		"results": [
             {
                 "name": "electric"
             }
         ]
	}
	`
}

// MockTypes mock
func MockTypes() PokemonTypes {
	return PokemonTypes{
		Types: []Type{{
			Name: "electric",
		}},
	}
}
