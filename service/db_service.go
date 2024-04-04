package service

import (
	"context"
	"database/sql"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"htmxx/db"
	"os"
)

type DBService struct {
}

func (s *DBService) Connect() (context.Context, *db.Queries, *sql.DB, error) {
	ctx := context.Background()
	url := os.Getenv("HTMXX_DB_URL")
	if url == "" {
		panic("HTMXX_DB_URL must be set")
	}
	dbConn, err := sql.Open("libsql", url)
	if err != nil {
		return nil, nil, nil, err
	}
	//defer dbConn.Close()
	queries := db.New(dbConn)
	return ctx, queries, dbConn, err
}
