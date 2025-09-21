package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./quotes.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to SQLite database")
	return db, nil
}
