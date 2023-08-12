package pokemon

import "errors"

var (
	ErrUnmarshalResponse = errors.New("can't unmarshal response")
	ErrPokemonNotFound   = errors.New("pokemon not found")
	ErrTypesNotFound     = errors.New("types not found")
	ErrCantPerformGet    = errors.New("can't perform get")
	ErrCantGetPokemon    = errors.New("can't get pokemon")
	ErrCantGetTypes      = errors.New("can't get types")
)
