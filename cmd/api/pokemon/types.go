package pokemon

import (
	"github.com/rromero96/PI-Pokemon/internal/pokemon"
)

// Pokemon contains the pokemon data
type Pokemon struct {
	ID      int
	Name    string
	HP      int
	Attack  int
	Defense int
	Image   string
	Speed   int
	Height  int
	Weight  int
	Created bool
	Types   []Type
}

// Type is part of pokemon
type Type struct {
	ID   int
	Name string
}

// toDTO converts a pokemon to a pokemonDTO
func (p Pokemon) toDTO() PokemonDTO {
	return PokemonDTO{
		ID:      &p.ID,
		Name:    &p.Name,
		HP:      &p.HP,
		Attack:  &p.Attack,
		Defense: &p.Defense,
		Image:   &p.Image,
		Speed:   &p.Speed,
		Height:  &p.Height,
		Weight:  &p.Weight,
		Created: &p.Created,
		Types:   toTypesDTO(p.Types),
	}
}

// toTypesDto converts a slice of type to a slice of typeDTO
func toTypesDTO(types []Type) []TypeDTO {
	typesDTO := make([]TypeDTO, len(types))
	for i, v := range types {
		typesDTO[i] = TypeDTO(v)
	}

	return typesDTO
}

// toTypesSlice converts a slice of pokemon.types to a slice of type
func toTypesSlice(pokemonTypes []pokemon.Type) []Type {
	types := make([]Type, len(pokemonTypes))
	for i, v := range pokemonTypes {
		types[i] = Type{
			ID:   i + 1,
			Name: v.Name,
		}
	}
	return types
}
