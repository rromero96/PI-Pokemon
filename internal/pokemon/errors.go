package pokemon

import "errors"

var (
	ErrUnmarshalResponse = errors.New("can't unmarshal response")
	ErrNotFound          = errors.New("pokemons not found")
)
