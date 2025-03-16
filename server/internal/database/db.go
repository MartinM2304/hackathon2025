package database

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/MartinM2304/hackathon2025/internal/models"

	_ "modernc.org/sqlite"
)

var db *sql.DB

var iteration int = 0

func InitDb() error {
	dbConn, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		return fmt.Errorf("Failed to open database: %v", err)
	}

	db = dbConn

	err = createTableIfNotExist()
	if err != nil {
		return fmt.Errorf("Failed to create table: %v", err)
	}

	return nil
}

func createTableIfNotExist() error {
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}

	slog.Info("Created the table")

	createTableSQL := `
	DROP TABLE IF EXISTS Votes;
	CREATE TABLE IF NOT EXISTS Votes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    type TEXT CHECK (type IN ('direction', 'emoji', 'sound')) NOT NULL,
    value INTEGER,
    turn INTEGER,
    timestamp DATETIME
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	return nil
}

func BatchInsertItems(items []models.DBser) error {
	if len(items) == 0 {
		return nil
	}

	err := db.Ping()
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	stmt, err := tx.Prepare("INSERT INTO Votes (type, value, turn, timestamp) VALUES (?, ?, ?, ?)")
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return fmt.Errorf("failed to rollback statement: %w", err)
		}
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for i, item := range items {
		dbDatum := item.DbData()
		_, err = stmt.Exec(dbDatum.DataType, dbDatum.Value, iteration, time.Now())
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				return fmt.Errorf("failed to rollback statement for item %d: %w", i, err)
			}
			return fmt.Errorf("failed to execute statement for item %d: %w", i, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit final transaction: %w", err)
	}
	iteration += 1
	return nil
}

func GetAllItemsForType(dataType string) (error, map[int][4]int) {
	result := map[int][4]int{}
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err), nil
	}

	query := `SELECT value, turn, COUNT(*) as value_count
		FROM Votes
		WHERE type = ?
		GROUP BY value
		ORDER BY value_count DESC;`

	rows, err := db.Query(query, dataType)
	if err != nil {
		return err, nil
	}
	defer rows.Close()

	for rows.Next() {
		var turn int
		var value int
		var count int
		if err := rows.Scan(&value, &turn, &count); err != nil {
			return fmt.Errorf("scan failed: %w", err), nil
		}
		arr, exists := result[turn]
		if !exists {
			arr = [4]int{}
		}

		arr[value]++

		result[turn] = arr
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows iteration error: %w", err), nil
	}

	return nil, result
}

func GetAllByType(dataType string) (error, []models.StatPair) {
	result := []models.StatPair{}
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err), nil
	}

	query := `SELECT value, COUNT(*) as count
		FROM Votes
		WHERE type = ?
		GROUP BY value
		ORDER BY count DESC`

	rows, err := db.Query(query, dataType)
	if err != nil {
		return err, nil
	}
	defer rows.Close()

	for rows.Next() {
		var value int
		var count int
		if err := rows.Scan(&value, &count); err != nil {
			return fmt.Errorf("scan failed: %w", err), nil
		}

		result = append(result, models.StatPair{
			Id:    value,
			Count: count,
		})

	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows iteration error: %w", err), nil
	}

	return nil, result
}

func DumpVotesTable() ([]models.Vote, error) {
	// SQL query to get all data from the Votes table
	query := `
		SELECT id, type, value, turn, timestamp
		FROM Votes
		ORDER BY id
	`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	// Collect the results
	var votes []models.Vote
	for rows.Next() {
		var v models.Vote
		var timestamp sql.NullString // Handle potential NULL values in timestamp

		if err := rows.Scan(&v.ID, &v.Type, &v.Value, &v.Turn, &timestamp); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		// Convert timestamp string to time.Time if not NULL
		if timestamp.Valid {
			// SQLite datetime format can vary, adjust the parsing format if needed
			t, err := time.Parse(time.RFC3339Nano, timestamp.String)
			if err != nil {
				// Try alternative format if the first one fails
				t, err = time.Parse(time.RFC3339Nano, timestamp.String)
				if err != nil {
					return nil, fmt.Errorf(
						"failed to parse timestamp '%s': %w",
						timestamp.String,
						err,
					)
				}
			}
			v.Timestamp = t
		}

		votes = append(votes, v)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return votes, nil
}

func CloseDb() error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
