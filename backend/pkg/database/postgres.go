package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDatabase() (*sql.DB, error) {
	host := "localhost" // Use "movie-chat-postgres" if your Go app is also in a Docker container
	port := 5433
	user := "postgres"
	password := "password"
	dbname := "go-chat"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
