package pokemon

const imageDefault string = "https://pokeapi.co/api/v2/pokemon/1/"

type PokemonDTO struct {
	ID      *int      `json:"id,omitempty"`
	Name    *string   `json:"name"`
	HP      *int      `json:"hp"`
	Attack  *int      `json:"attack"`
	Defense *int      `json:"defense"`
	Image   *string   `json:"image,,omitempty"`
	Speed   *int      `json:"speed"`
	Height  *int      `json:"height"`
	Weight  *int      `json:"weight"`
	Created *bool     `json:"created"`
	Types   []TypeDTO `json:"types"`
}

type TypeDTO struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func (p PokemonDTO) validate() error {
	if p.Name == nil || p.HP == nil || p.Attack == nil || p.Defense == nil || p.Speed == nil || p.Height == nil || p.Weight == nil {
		return ErrInvalidBody
	}

	if p.Types == nil || len(p.Types) == 0 {
		return ErrInvalidBody
	}

	for _, v := range p.Types {
		return v.validate()
	}

	return nil
}

func (t TypeDTO) validate() error {
	if t.Name == "" {
		return ErrInvalidBody
	}

	return nil
}

func (p PokemonDTO) toDomain() Pokemon {
	id := 0
	if p.ID != nil {
		id = *p.ID
	}

	image := imageDefault
	if p.Image != nil {
		image = *p.Image
	}

	created := false
	if p.Created != nil {
		created = *p.Created
	}

	return Pokemon{
		ID:      id,
		Name:    *p.Name,
		HP:      *p.HP,
		Attack:  *p.Attack,
		Defense: *p.Defense,
		Image:   image,
		Speed:   *p.Speed,
		Height:  *p.Height,
		Weight:  *p.Weight,
		Created: created,
		Types:   toTypes(p.Types),
	}
}

func toTypes(typesDTO []TypeDTO) []Type {
	types := make([]Type, len(typesDTO))
	for i, v := range typesDTO {
		types[i] = Type{
			Name: v.Name,
		}
	}

	return types
}
