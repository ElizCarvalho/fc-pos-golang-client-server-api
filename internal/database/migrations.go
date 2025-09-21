package database

import (
	"database/sql"
	"log"
)

func CreateTable(db *sql.DB) error {

	query := `
	CREATE TABLE IF NOT EXISTS quotes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid REAL NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("Table 'quotes' created/verified successfully")
	return nil
}
