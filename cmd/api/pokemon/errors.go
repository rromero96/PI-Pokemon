package pokemon

import "errors"

var (
	ErrCantPrepareStatement = errors.New("can't prepare statement")
	ErrCantRunQuery         = errors.New("can't run query")
	ErrCantScanRowResult    = errors.New("can't scan row result")
	ErrCantReadRows         = errors.New("can't read rows")
)
