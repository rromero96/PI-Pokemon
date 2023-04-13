package pokemon

import "github.com/rromero96/PI-Pokemon/internal/pokemon"

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

type Type struct {
	ID   int
	Name string
}

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

func toTypesDTO(types []Type) []TypeDTO {
	typesDTO := make([]TypeDTO, len(types))
	for i, v := range types {
		typesDTO[i] = TypeDTO{
			Name: &v.Name,
		}
	}

	return typesDTO
}

func toTypesSlice(pokemonTypes []pokemon.Type) []Type {
	types := make([]Type, len(pokemonTypes))
	for i, v := range pokemonTypes {
		types[i] = Type{
			Name: v.Name,
		}
	}

	return types
}
