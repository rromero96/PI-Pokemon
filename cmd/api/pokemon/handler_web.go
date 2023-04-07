package pokemon

import (
	"net/http"

	"github.com/rromero96/roro-lib/cmd/web"
)

// SearchV1 performs a search to obtain all the pokemons
func SearchV1() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

// SearchVByIDV1 performs a search to obtain a pokemon by ID
func SearchByIDV1() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

// CreateV1 perfoms a pokemon creation
func CreateV1() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

// SearchTypesV1 performs a search to obtain all pokemon types
func SearchTypesV1() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}
