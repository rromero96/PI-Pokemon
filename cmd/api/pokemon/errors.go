package pokemon

import "errors"

var (
	ErrInvalidBody          = errors.New(InvalidBody)
	ErrCantPrepareStatement = errors.New("can't prepare statement")
	ErrCantRunQuery         = errors.New("can't run query")
	ErrCantScanRowResult    = errors.New("can't scan row result")
	ErrCantReadRows         = errors.New("can't read rows")
	ErrCantAddTypes         = errors.New("can't add types")
	ErrCantGetLastID        = errors.New("can't get last id")
	ErrCantGetTypes         = errors.New("can't get types")
	ErrCantGetApiTypes      = errors.New("can't get api types")
	ErrCantGetPokemon       = errors.New("can't get pokemon")
	ErrCantGetApiPokemon    = errors.New("can't get api Pokemon")
	ErrCantCreatePokemon    = errors.New(CantCreatePokemon)
	ErrCantSaveTypes        = errors.New("can't save types")
	ErrPokemonNotFound      = errors.New("pokemon not found")
	ErrInvalidPokemon       = errors.New(InvalidPokemon)
	ErrCantBeginTransaction = errors.New("can't begin transaction")
)

const (
	InvalidBody       string = "invalid body"
	InvalidPokemon    string = "invalid pokemon"
	InvalidID         string = "invalid id"
	CantCreatePokemon string = "can't create pokemon"
	CantSearchPokemon string = "can't search pokemon"
	CantGetTypes      string = "can't get types'"
	NotFound          string = "pokemon not found"
)
