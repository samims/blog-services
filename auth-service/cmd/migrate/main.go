package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"auth-service/constants"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var migrationDir string
	flag.StringVar(&migrationDir, "migration-dir", "migrations", "Directory with migration files")
	flag.Parse()

	dbURL := getConnectionURL()
	if dbURL == "" {
		log.Fatal("DATABASE_URL generation failed")
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationDir),
		dbURL,
	)
	if err != nil {
		log.Fatal(err)
	}

	if migrationErr := m.Up(); migrationErr != nil && !errors.Is(migrationErr, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	log.Println("Migrations completed successfully")

	// Print tables after migration
	if err := printTables(dbURL); err != nil {
		log.Printf("Error printing tables: %v", err)
	}
}

func getConnectionURL() string {
	dbHost := os.Getenv(constants.PostgresHost)
	dbPort := os.Getenv(constants.PostgresPort)
	dbUser := os.Getenv(constants.PostgresUser)
	dbPass := os.Getenv(constants.PostgresPasswd)
	dbName := os.Getenv(constants.PostgresDBName)
	dbSSLMode := os.Getenv(constants.DatabaseSSLMode)

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		dbSSLMode,
	)
	return dbURL

}

func printTables(dbURL string) error {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error closing database connection: %v", err)
		}
	}(db)

	rows, err := db.Query(
		"SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_name = 'BASE TABLE'")
	if err != nil {
		return fmt.Errorf("error querying tables: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("error closing rows: %v", err)
		}
	}(rows)

	log.Println("Tables in the database:")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return fmt.Errorf("error scanning table name: %w", err)
		}
		log.Printf("- %s", tableName)

		// Print columns for each table
		if err := printColumns(db, tableName); err != nil {
			log.Printf("  Error printing columns for %s: %v", tableName, err)
		}
	}

	return rows.Err()
}

func printColumns(db *sql.DB, tableName string) error {
	rows, err := db.Query(
		"SELECT column_name, data_type FROM information_schema.columns WHERE table_name = $1", tableName)
	if err != nil {
		return fmt.Errorf("error querying columns: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}(rows)

	log.Printf("  Columns:")
	for rows.Next() {
		var columnName, dataType string
		if err := rows.Scan(&columnName, &dataType); err != nil {
			return fmt.Errorf("error scanning column info: %w", err)
		}
		log.Printf("    - %s (%s)", columnName, dataType)
	}

	return rows.Err()
}
