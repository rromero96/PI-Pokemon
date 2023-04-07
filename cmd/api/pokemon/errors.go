package pokemon

import "errors"

var (
	ErrInvalidBody          = errors.New(InvalidBody)
	ErrCantPrepareStatement = errors.New("can't prepare statement")
	ErrCantRunQuery         = errors.New("can't run query")
	ErrCantScanRowResult    = errors.New("can't scan row result")
	ErrCantReadRows         = errors.New("can't read rows")
)

const (
	InvalidBody       string = "invalid body"
	BadRequest        string = "bad request"
	CantCreatePokemon string = "can't create pokemon"
)
