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

func CloseDb() error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
