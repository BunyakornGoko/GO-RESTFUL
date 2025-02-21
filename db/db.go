package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	dbFile := os.Getenv("DB_CONNECTION_STRING")
	if dbFile == "" {
		dbFile = "api.db"
	}
	DB, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		return err
	}

	err = createTables()
	if err != nil {
		return err
	}

	return nil
}

func createTables() error {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`
	
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		return err
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		return err
	}

	createRegistrationsTables := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createRegistrationsTables)
	if err != nil {
		return err
	}

	return nil
}