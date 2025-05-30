package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB // DB is a global variable that holds the database connection

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite", "api.db") // Open a connection to the SQLite database named "api.db"
	// If the database does not exist, it will be created automatically

	if err != nil {
		panic("Could not connect to the database.")
	}

	DB.SetMaxOpenConns(10) // Set the maximum number of open connections to the database
	DB.SetMaxIdleConns(5)  // Set the maximum number of idle connections in the pool

	createTables() // Call the function to create the necessary tables in the database
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable) // Execute the SQL command to create the users table if it does not exist

	if err != nil {
		panic("Could not create users table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME  NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable) // Execute the SQL command to create the events table if it does not exist

	if err != nil {
		panic("Could not create events table.")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id)
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable) // Execute the SQL command to create the registrations table if it does not exist

	if err != nil {
		panic("Could not create registrations table.")
	}
}
