package db

import (
	"database/sql"

	"example.com/rest-api/helper"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	helper.PanicError(err, "Could not connect to database")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTables := `
	CREATE TABLE IF NOT EXISTS events (
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 name TEXT NOT NULL,
	 description TEXT NOT NULL,
	 location TEXT NOT NULL,
	 date DATETIME NOT NULL,
	 user_id INTEGER 
	)
	`

	_, err := DB.Exec(createEventsTables)

	helper.PanicError(err, "Could not create events table")
}
