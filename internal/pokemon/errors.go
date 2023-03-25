package pokemon

import "errors"

var (
	ErrUnmarshalResponse = errors.New("can't unmarshal response")
	ErrPokemonNotFound   = errors.New("pokemon not found")
	ErrTypesNotFound     = errors.New("types not found")
)
