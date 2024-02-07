package pokemon

import (
	"context"
	"database/sql"
	"fmt"
)

// MockMySQLCreate mock
func MockMySQLCreate(err error) MySQLCreate {
	return func(context.Context, Pokemon) error {
		return err
	}
}

// MockMySQLSearchByID mock
func MockMySQLSearchByID(pokemon Pokemon, err error) MySQLSearchByID {
	return func(context.Context, int) (Pokemon, error) {
		return pokemon, err
	}
}

// MockMySQLCreate mock
func MockMySQLAdd(err error) MySQLAdd {
	return func(context.Context, int, []Type, *sql.Tx) error {
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

// MockSearchByID mock
func MockSearchByID(pokemon Pokemon, err error) SearchByID {
	return func(ctx context.Context, ID int) (Pokemon, error) {
		return pokemon, err
	}
}

// MockCreate mock
func MockCreate(err error) Create {
	return func(ctx context.Context, pokemon Pokemon) error {
		return err
	}
}

// MockPokemonAsJson mock
func MockPokemonAsJson() string {
	return fmt.Sprintf(`
	{
		"id": 1,
		"name": "bulbasaur",
		"hp": 100,
		"attack": 100,
		"defense": 100,
		"image": "image",
		"speed": 100,
		"height": 100,
		"weight": 100,
		"types" : %v
	}
	`, MockTypesAsJson())
}

// MockPokemon mock
func MockPokemon() Pokemon {
	return Pokemon{
		ID:      1,
		Name:    "bulbasaur",
		HP:      100,
		Attack:  100,
		Defense: 100,
		Image:   "image",
		Speed:   100,
		Height:  100,
		Weight:  100,
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
			Name: "grass",
		},
		{
			Name: "poison",
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
			"name": "grass"
		},
		{
			"name": "poison"
		}
	]
	`
}
