package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rromero96/roro-lib/cmd/config"
	"github.com/rromero96/roro-lib/cmd/rest"
	"github.com/rromero96/roro-lib/cmd/web"

	"github.com/rromero96/PI-Pokemon/cmd/api/pokemon"
	internalPokemon "github.com/rromero96/PI-Pokemon/internal/pokemon"
)

const (
	pokemonsSearchV1      string = "/pokemons/v1"
	pokemonsSearchTypesV1 string = "/pokemons/types/v1"
	pokemonCreateV1       string = "/pokemon/v1"
	pokemonSearchByIDV1   string = "/pokemon/{pokemon_id}/v1"

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
	/*	mysql	*/
	addTypes := pokemon.MakeMySQLAdd(pokemonsDBClient)
	mysqlCreatePokemon := pokemon.MakeMySQLCreate(pokemonsDBClient, addTypes)
	mysqlCreateTypes := pokemon.MakeMySQLCreateType(pokemonsDBClient)
	mysqlSearchTypes, err := pokemon.MakeMySQLSearchTypes(pokemonsDBClient)
	if err != nil {
		return err
	}

	/*	internal	*/
	restGetFunc := rest.MakeGetFunc("", "")
	pokemonTypeSearch, err := internalPokemon.MakeSearchTypes(restGetFunc)
	if err != nil {
		return err
	}

	/*	api	*/
	searchTypes := pokemon.MakeSearchTypes(mysqlSearchTypes, pokemonTypeSearch, mysqlCreateTypes)

	/*
		Endpoints
	*/
	app.Get(pokemonsSearchV1, pokemon.SearchV1())
	app.Get(pokemonsSearchTypesV1, pokemon.SearchTypesV1(searchTypes))

	app.Post(pokemonCreateV1, pokemon.CreateV1(mysqlCreatePokemon))
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
	dbUsername := config.String("databases", fmt.Sprintf("mysql.%s.username", database), "")
	dbPassword := config.String("databases", fmt.Sprintf("mysql.%s.password", database), "")
	dbName := config.String("databases", fmt.Sprintf("mysql.%s.db_name", database), "")
	return fmt.Sprintf(connectionStringFormat, dbUsername, dbPassword, dbName)
}
