package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rromero96/roro-lib/cmd/web"

	"github.com/rromero96/PI-Pokemon/cmd/api/pokemon"
)

const (
	pokemonsSearchV1      string = "/pokemons"
	pokemonsSearchTypesV1 string = "/pokemons/types"
	pokemonCreateV1       string = "/pokemon"
	pokemonSearchByIDV1   string = "/pokemon/{pokemon_id}"

	// connectionStringFormat when its deployed needs to have the host next to @tcp, check https://github.com/go-sql-driver/mysql/
	connectionStringFormat string = "%s:%s@tcp/%s?charset=utf8&parseTime=true"
	mysqlDriver            string = "mysql"
	pokemonsDB             string = "pokemons"
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
	   MYSQL client
	*/
	pokemonsDBClient, err := createDBClient(getDBConnectionStringRoutes(pokemonsDB))
	if err != nil {
		return err
	}

	/*
		Injections
	*/
	mysqlCreatePokemonFunc := pokemon.MakeMySQLCreate(pokemonsDBClient)
	if err != nil {
		panic(err)
	}

	/*
		Endpoints
	*/
	app.Get(pokemonsSearchV1, pokemon.SearchV1())
	app.Get(pokemonsSearchTypesV1, pokemon.SearchTypesV1())

	app.Post(pokemonCreateV1, pokemon.CreateV1(mysqlCreatePokemonFunc))
	app.Get(pokemonSearchByIDV1, pokemon.SearchByIDV1())

	log.Print("server up and running in port 8080")
	return web.Run(ln, web.DefaultTimeouts, app)
}

func createDBClient(connectionString string) (*sql.DB, error) {
	db, err := sql.Open(mysqlDriver, connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(14 * time.Minute)

	return db, nil
}

func getDBConnectionStringRoutes(database string) string {
	/* 	dbUsername := config.String("databases", fmt.Sprintf("mysql.%s.username", database), "")
	   	dbPassword := ""
	   	dbName := config.String("databases", fmt.Sprintf("mysql.%s.db_name", database), "") */
	dbUsername := "root"
	dbPassword := ""
	dbName := "pokemons"
	return fmt.Sprintf(connectionStringFormat, dbUsername, dbPassword, dbName)
}
