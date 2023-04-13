package pokemon

import (
	"context"
	"fmt"
)

// MockMySQLCreate mock
func MockMySQLCreate(err error) MySQLCreate {
	return func(context.Context, Pokemon) error {
		return err
	}
}

// MockMySQLCreate mock
func MockMySQLAdd(err error) MySQLAdd {
	return func(context.Context, int, []Type) error {
		return err
	}
}

// MockMySQLCreateType mock
func MockMySQLCreateType(err error) MySQLCreateType {
	return func(context.Context, []Type) error {
		return err
	}
}

// MockMySQLSearchTypes mock
func MockMySQLSearchTypes(res []Type, err error) MySQLSearchTypes {
	return func(context.Context) ([]Type, error) {
		return res, err
	}
}

// MockSearchTypes mock
func MockSearchTypes(res []Type, err error) SearchTypes {
	return func(ctx context.Context) ([]Type, error) {
		return res, err
	}
}

// MockPokemonAsJson mock
func MockPokemonAsJson() string {
	return fmt.Sprintf(`
	{
		"id": 25,
		"name": "pikachu",
		"hp": 35,
		"attack": 55,
		"defense": 40,
		"image": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/dream-world/25.svg",
		"speed": 90,
		"height": 4,
		"weight": 60,
		"types" : %v
	}
	`, MockTypesAsJson())
}

// MockPokemon mock
func MockPokemon() Pokemon {
	return Pokemon{
		ID:      25,
		Name:    "pikachu",
		HP:      35,
		Attack:  55,
		Defense: 40,
		Image:   "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/dream-world/25.svg",
		Speed:   90,
		Height:  4,
		Weight:  60,
		Types:   MockTypes(),
	}
}

// MockPokemonDTO mock
func MockPokemonDTO() PokemonDTO {
	return MockPokemon().toDTO()
}

// MockTypes  mock
func MockTypes() []Type {
	return []Type{
		{
			Name: "electric",
		},
	}
}

// MockTypesDTO
func MockTypesDTO() []TypeDTO {
	return toTypesDTO(MockTypes())
}

// MockTypesAsJson mock
func MockTypesAsJson() string {
	return `
	[
		{
			"name": "electric"
		}
	]
	`
}
