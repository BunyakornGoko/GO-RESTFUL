package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("mysql", "root:!Zazaza080a@tcp(127.0.0.1:3306)/events")
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
		id INT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	)`
	
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		return err
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		location VARCHAR(255) NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INT,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		return err
	}

	createRegistrationsTables := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INT AUTO_INCREMENT PRIMARY KEY,
		event_id INT,
		user_id INT,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createRegistrationsTables)
	if err != nil {
		return err
	}

	return nil
}