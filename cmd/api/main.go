package main

import (
	"log"
	"net"
	"net/http"

	"github.com/rromero96/roro-lib/cmd/web"
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
	app.Get("/", Hello())

	log.Print("server up and running in port 8080")
	return web.Run(ln, web.DefaultTimeouts, app)
}

func Hello() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return web.EncodeJSON(w, []byte("server upp"), http.StatusOK)
	}
}
