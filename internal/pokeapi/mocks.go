package pokeapi

import (
	"context"
)

// MockGetByID mock
func MockGetByID(response PokemonDTO, err error) GetByID {
	return func(context.Context, int) (PokemonDTO, error) {
		return response, err
	}
}

// MockGetTypes mock
func MockGetTypes(response PokemonTypesDTO, err error) GetTypes {
	return func(context.Context) (PokemonTypesDTO, error) {
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
        },
		{
            "base_stat": 100,
            "stat": {
                "name": "attack"
            }
        },
		{
            "base_stat": 100,
            "stat": {
                "name": "defense"
            }
        },
		{
            "base_stat": 100,
            "stat": {
                "name": "special-attack"
            }
        },
		{
            "base_stat": 100,
            "stat": {
                "name": "special-defense"
            }
        },
		{
            "base_stat": 100,
            "stat": {
                "name": "speed"
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

// MockPokemonDTO mock
func MockPokemonDTO() PokemonDTO {
	return PokemonDTO{
		ID:     1,
		Name:   "bulbasaur",
		Height: 100,
		Weight: 100,
		Sprites: SpritesDTO{
			Other: OtherDTO{
				DreamWorld: DreamWorldDTO{
					FrontDefault: "image",
				},
			},
		},
		Stats: []StatsDTO{
			{
				BaseStat: 100,
				Stat:     StatDTO{Name: "hp"},
			},
			{
				BaseStat: 100,
				Stat:     StatDTO{Name: "attack"},
			},
			{
				BaseStat: 100,
				Stat:     StatDTO{Name: "defense"},
			},
			{
				BaseStat: 100,
				Stat:     StatDTO{Name: "special-attack"},
			},
			{
				BaseStat: 100,
				Stat:     StatDTO{Name: "special-defense"},
			},
			{
				BaseStat: 100,
				Stat:     StatDTO{Name: "speed"},
			},
		},
		Types: []TypesDTO{{
			Type: TypeDTO{Name: "grass"},
		}, {
			Type: TypeDTO{Name: "poison"},
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

// MockPodemonTypesDTO mock
func MockPokemonTypesDTO() PokemonTypesDTO {
	return PokemonTypesDTO{
		Types: []TypeDTO{{
			Name: "grass",
		}, {
			Name: "poison",
		}},
	}
}
