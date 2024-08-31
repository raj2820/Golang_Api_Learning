package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB establishes a connection to the SQLite database "api.db".
func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}

	// Set connection pool configuration (optional)
	 DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	err = createTables()
	if err != nil {
		return fmt.Errorf("error creating tables: %w", err)
	}

	return nil
}

// createTables defines and executes the SQL statements to create the `events` table.
func createTables() error {
	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime DATETIME NOT NULL,
        user_id INTEGER
    );
    `

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		return fmt.Errorf("error creating events table: %w", err)
	}

	return nil
}
