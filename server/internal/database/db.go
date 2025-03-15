package database

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	"github.com/MartinM2304/hackathon2025/internal/models"

	_ "modernc.org/sqlite"
)

var db *sql.DB

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
	CREATE TABLE IF NOT EXISTS Data (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Type TEXT NOT NULL,
		Value INTEGER,
		IpAddress TEXT
	)`

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

	fmt.Println(db)
	// err := db.Ping()
	// if err != nil {
	// 	return fmt.Errorf("Failed to connect to database: %v", err)
	// }

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	stmt, err := tx.Prepare("INSERT INTO Data (Type, Value, IpAddress) VALUES (?, ?, ?)")
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
		_, err = stmt.Exec(dbDatum.DataType, dbDatum.Value, dbDatum.IpAddr)
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

	return nil
}

func CloseDb() error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
