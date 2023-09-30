package core

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mitchellh/go-homedir"
)

const dbName = "ratings.db"

func WithDB(callback func(db *sql.DB) error) error {
	db, err := setupDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Call the provided callback function with the database connection.
	return callback(db)
}

func setupDatabase() (*sql.DB, error) {
	// Specify the database path.
	appName := "mood-tracker"
	dbPath := filepath.Join("~", ".cache", appName, dbName)

	// Expand the tilde (~) in the path.
	expandedDBPath, err := homedir.Expand(dbPath)
	if err != nil {
		return nil, err
	}

	// Create the directory structure if it doesn't exist.
	dir := filepath.Dir(expandedDBPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, err
	}

	// Open a connection to the SQLite database.
	db, err := sql.Open("sqlite3", expandedDBPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}
