package pokemon

type (
	Pokemon struct {
		ID      int     `json:"id"`
		Name    string  `json:"name"`
		Height  int     `json:"height"`
		Weight  int     `json:"weight"`
		Sprites Sprites `json:"sprites"`
		Stats   []Stats `json:"stats"`
		Types   []Types `json:"types"`
	}

	PokemonTypes struct {
		Types []Type `json:"results"`
	}

	Stats struct {
		BaseStat int  `json:"base_stat"`
		Stat     Stat `json:"stat"`
	}

	Stat struct {
		Name string `json:"name"`
	}

	Types struct {
		Type Type `json:"type"`
	}

	Type struct {
		Name string `json:"name"`
	}

	Sprites struct {
		Other Other `json:"other"`
	}

	Other struct {
		DreamWorld DreamWorld `json:"dream_world"`
	}

	DreamWorld struct {
		FrontDefault string `json:"front_default"`
	}
)
