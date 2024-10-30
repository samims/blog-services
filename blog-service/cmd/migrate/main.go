package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"

	"blog-service/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	migrationDir := flag.String("migration-dir", "migrations", "Directory with migration files")
	flag.Parse()
	cfg := config.NewPostgresConfig(viper.New())

	dbURL := cfg.ConnectionURL()

	m, err := migrate.New(fmt.Sprintf("file://%s", *migrationDir), dbURL)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Println("Migrations completed successfully")

	if err := printDatabaseSchema(dbURL); err != nil {
		log.Printf("Warning: Failed to print database schema: %v", err)
	}

	return nil
}

func printDatabaseSchema(dbURL string) error {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	tables, err := getTables(db)
	if err != nil {
		return fmt.Errorf("failed to get tables: %w", err)
	}

	log.Println("Database schema:")
	for _, table := range tables {
		log.Printf("- Table: %s", table)
		columns, err := getColumns(db, table)
		if err != nil {
			log.Printf("  Warning: Failed to get columns for %s: %v", table, err)
			continue
		}
		for _, column := range columns {
			log.Printf("  - %s (%s)", column.Name, column.Type)
		}
	}

	return nil
}

func getTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`
		SELECT table_name 
		FROM information_schema.tables 
		WHERE table_schema = 'public' AND table_type = 'BASE TABLE'
	`)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf(" Waring: Failed to close rows: %v", err)
			return
		}
	}(rows)

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}
	return tables, rows.Err()
}

type column struct {
	Name string
	Type string
}

func getColumns(db *sql.DB, tableName string) ([]column, error) {
	rows, err := db.Query(`
		SELECT column_name, data_type 
		FROM information_schema.columns 
		WHERE table_schema = 'public' AND table_name = $1
	`, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []column
	for rows.Next() {
		var col column
		if err := rows.Scan(&col.Name, &col.Type); err != nil {
			return nil, err
		}
		columns = append(columns, col)
	}
	return columns, rows.Err()
}
