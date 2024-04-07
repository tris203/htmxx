package service

import (
	"database/sql"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"htmxx/db"
	"os"
)

type DBService struct {
}

func (s *DBService) Connect() (*db.Queries, *sql.DB, error) {
	url := os.Getenv("HTMXX_DB_URL")
	if url == "" {
		panic("HTMXX_DB_URL must be set")
	}
	dbConn, err := sql.Open("libsql", url)
	if err != nil {
		return nil, nil, err
	}
	//defer dbConn.Close()
	queries := db.New(dbConn)
	return queries, dbConn, err
}
