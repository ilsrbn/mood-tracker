package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"mood-tracker/src/core"
	"time"
)

type Rating struct {
	ID        int
	Value     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Call the provided callback function with the database connection.
func CreateRatingTable(db *sql.DB) error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS Rating (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			value INTEGER,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		)
	`
	_, err := db.Exec(createTableQuery)
	return err
}

func InsertRating(value int) (int64, error) {
	var lastInsertID int64

	// Define the database callback function.
	callback := func(db *sql.DB) error {
		insertRatingQuery := `
            INSERT INTO Rating (value, created_at, updated_at) VALUES (?, ?, ?)
        `
		currentTime := time.Now()
		result, err := db.Exec(insertRatingQuery, value, currentTime, currentTime)
		if err != nil {
			return err
		}
		lastInsertID, err = result.LastInsertId()
		return err
	}

	// Use the WithDB function from the database package to execute the callback.
	if err := core.WithDB(callback); err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func GetAllRatings() ([]Rating, error) {
	var ratings []Rating

	// Define the database callback function.
	callback := func(db *sql.DB) error {
		queryRatingsQuery := `
			SELECT * FROM Rating
		`
		rows, err := db.Query(queryRatingsQuery)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var rating Rating
			err := rows.Scan(&rating.ID, &rating.Value, &rating.CreatedAt, &rating.UpdatedAt)
			if err != nil {
				return err
			}
			ratings = append(ratings, rating)
		}

		if err := rows.Err(); err != nil {
			return err
		}

		return nil
	}

	// Use the WithDB function from the database package to execute the callback.
	if err := core.WithDB(callback); err != nil {
		return nil, err
	}

	return ratings, nil
}

// UpdateRating updates a rating by ID.
func UpdateRating(id int, value int) error {
	// Define the database callback function.
	callback := func(db *sql.DB) error {
		updateRatingQuery := `
			UPDATE Rating SET value = ?, updated_at = ? WHERE id = ?
		`
		_, err := db.Exec(updateRatingQuery, value, time.Now(), id)
		return err
	}

	// Use the WithDB function from the database package to execute the callback.
	if err := core.WithDB(callback); err != nil {
		return err
	}

	return nil
}
