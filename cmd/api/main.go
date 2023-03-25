package main

import (
	"log"
	"net"

	"github.com/rromero96/roro-lib/cmd/web"

	"github.com/rromero96/PI-Pokemon/cmd/api/pokemon"
)

const (
	pokemonsSearchV1      string = "/pokemons"
	pokemonsSearchTypesV1 string = "/pokemons/types"
	pokemonCreateV1       string = "/pokemon"
	pokemonSearchByIDV1   string = "/pokemon/{pokemon_id}"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	/*
		Server Configuration
	*/
	app := web.New()

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	/*
		Endpoints
	*/
	app.Get(pokemonsSearchV1, pokemon.SearchV1())
	app.Get(pokemonsSearchTypesV1, pokemon.SearchTypesV1())

	app.Post(pokemonCreateV1, pokemon.CreateV1())
	app.Get(pokemonSearchByIDV1, pokemon.SearchByIDV1())

	log.Print("server up and running in port 8080")
	return web.Run(ln, web.DefaultTimeouts, app)
}
