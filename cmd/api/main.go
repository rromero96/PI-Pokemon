package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/olebedev/config"
	"github.com/rromero96/roro-lib/httpclient"
	"github.com/rromero96/roro-lib/web"

	"github.com/rromero96/PI-Pokemon/cmd/api/pokemon"
	"github.com/rromero96/PI-Pokemon/internal/pokeapi"
)

const (
	pokemonsGetV1      string = "/pokemons/v1"
	pokemonsGetTypesV1 string = "/pokemons/types/v1"
	pokemonCreateV1    string = "/pokemon/v1"
	pokemonGetByIDV1   string = "/pokemon/id/{pokemon_id}/v1"

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
	   YML Configuration
	*/
	file, err := os.ReadFile(getFileName("../PI-Pokemon/conf", "production.yml"))
	if err != nil {
		panic(err)
	}
	yamlString := string(file)

	cfg, _ := config.ParseYaml(yamlString)

	/*
	   MYSQL client
	*/
	pokemonsDBClient, err := createDBClient(getDBConnectionStringRoutes(pokemonsDB, cfg))
	if err != nil {
		return err
	}

	/*
	   HTTP client
	*/
	connectionTimeout := 5000 * time.Millisecond
	retries := 1
	opts := []httpclient.OptionRetryable{
		httpclient.WithTimeout(connectionTimeout),
	}
	httpClient := httpclient.NewRetryable(retries, opts...)

	/*
		Injections
	*/

	/*	mysql	*/
	addTypes := pokemon.MakeMySQLAdd(pokemonsDBClient)
	mysqlCreate := pokemon.MakeMySQLCreate(pokemonsDBClient, addTypes)
	mysqlSearchByID := pokemon.MakeMySQLSearchByID(pokemonsDBClient)
	mysqlCreateTypes := pokemon.MakeMySQLCreateType(pokemonsDBClient)
	mysqlSearchTypes := pokemon.MakeMySQLSearchTypes(pokemonsDBClient)

	/*	internal	*/
	getTypes, err := pokeapi.MakeGetTypes(httpClient)
	if err != nil {
		return err
	}

	getByID, err := pokeapi.MakeGetByID(httpClient)
	if err != nil {
		return err
	}

	/*	api	*/
	searchTypes := pokemon.MakeGetTypes(mysqlSearchTypes, getTypes, mysqlCreateTypes)
	searchByID := pokemon.MakeGetByID(mysqlSearchByID, getByID, mysqlCreate)
	create := pokemon.MakeCreate(mysqlCreate)

	/*
		Endpoints
	*/
	app.Get(pokemonsGetV1, pokemon.GetV1())
	app.Get(pokemonsGetTypesV1, pokemon.GetTypesV1(searchTypes))

	app.Post(pokemonCreateV1, pokemon.CreateV1(create))
	app.Get(pokemonGetByIDV1, pokemon.GetByIDV1(searchByID))

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

func getDBConnectionStringRoutes(database string, yml *config.Config) string {
	dbUserName, _ := yml.String(fmt.Sprintf("databases.mysql.%s.username", database))
	dbPassword, _ := yml.String(fmt.Sprintf("databases.mysql.%s.password", database))
	//dbHost, _ := yml.String(fmt.Sprintf("databases.mysql.%s.db_host", database))
	dbName, _ := yml.String(fmt.Sprintf("databases.mysql.%s.db_name", database))
	return fmt.Sprintf(connectionStringFormat, dbUserName, dbPassword /* dbHost ,*/, dbName)
}

// getFileName returns the absolute file path of a file
func getFileName(folder string, file string) string {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	rootDir := filepath.Join(currentDir, "..", "..")

	return filepath.Join(rootDir, folder, file)
}
