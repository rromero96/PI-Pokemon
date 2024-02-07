package pokeapi

type (
	PokemonDTO struct {
		ID      int        `json:"id"`
		Name    string     `json:"name"`
		Height  int        `json:"height"`
		Weight  int        `json:"weight"`
		Sprites SpritesDTO `json:"sprites"`
		Stats   []StatsDTO `json:"stats"`
		Types   []TypesDTO `json:"types"`
	}

	PokemonTypesDTO struct {
		Types []TypeDTO `json:"results"`
	}

	StatsDTO struct {
		BaseStat int     `json:"base_stat"`
		Stat     StatDTO `json:"stat"`
	}

	StatDTO struct {
		Name string `json:"name"`
	}

	TypesDTO struct {
		Type TypeDTO `json:"type"`
	}

	TypeDTO struct {
		Name string `json:"name"`
	}

	SpritesDTO struct {
		Other OtherDTO `json:"other"`
	}

	OtherDTO struct {
		DreamWorld DreamWorldDTO `json:"dream_world"`
	}

	DreamWorldDTO struct {
		FrontDefault string `json:"front_default"`
	}
)
