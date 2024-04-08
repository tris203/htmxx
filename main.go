package main

import (
	"database/sql"
	"fmt"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"htmxx/db"
	"htmxx/service"
	"net/http"
	"os"
)

type application struct {
	config        config
	router        *http.ServeMux
	db            *sql.DB
	query         *db.Queries
	eventsService *service.EventsService
}

type config struct {
	httpPort string
	db       struct {
		dsn string
	}
}

func main() {

	var config config

	config.httpPort = GetStringEnv("PORT", "8081")
	config.db.dsn = GetStringEnv("HTMXX_DB_URL", "db.sqlite")

	dbConn, err := NewDB(config.db.dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dbConn.Close()

	app := &application{
		config: config,
		db:     dbConn,
		query:  db.New(dbConn),
	}

	app.serveHTTP()

}

func GetStringEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func NewDB(dsn string) (*sql.DB, error) {

	dbConn, err := sql.Open("libsql", dsn)
	if err != nil {
		return nil, err
	}
	return dbConn, err
}
