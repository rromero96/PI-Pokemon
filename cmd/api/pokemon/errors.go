package pokemon

import "errors"

var (
	ErrInvalidBody            = errors.New(InvalidBody)
	ErrCantPrepareStatement   = errors.New("can't prepare statement")
	ErrCantRunQuery           = errors.New("can't run query")
	ErrCantScanRowResult      = errors.New("can't scan row result")
	ErrCantReadRows           = errors.New("can't read rows")
	ErrCantAddTypes           = errors.New("can't add types")
	ErrCantGetLastID          = errors.New("can't get last id")
	ErrCantSearchTypes        = errors.New("can't search types")
	ErrCantSearchPokemonTypes = errors.New("can't search pokemon types")
	ErrCantSaveTypes          = errors.New("can't save types")
)

const (
	InvalidBody       string = "invalid body"
	InvalidPokemon    string = "invalid pokemon"
	CantCreatePokemon string = "can't create pokemon"
	CantGetTypes      string = "can't get types'"
)
