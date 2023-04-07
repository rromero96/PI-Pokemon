package pokemon

type PokemonDTO struct {
	ID      *string   `json:"id,omitempty"`
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
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name"`
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
	if t.Name == nil {
		return ErrInvalidBody
	}

	return nil
}
