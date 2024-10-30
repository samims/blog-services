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
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	l := logrus.StandardLogger()
	l.Info("----------------------")
	l.Infof("starting migrations.....")
	l.Info("----------------------")

	if err := run(); err != nil {
		l.Fatalf("Error: %v", err)
	}
}

func run() error {
	migrationDir := flag.String("migration-dir", "migrations", "Directory with migration files")
	flag.Parse()
	cfg := config.NewPostgresConfig(viper.New())

	dbURL := cfg.ConnectionURLWithScheme()
	// set scheme to dbURL

	m, err := migrate.New(fmt.Sprintf("file://%s", *migrationDir), dbURL)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	log.Println("################\nmigrating...\n#####################")

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Println("Migrations completed successfully...")

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
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}(db)

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
		SELECT columns.column_name, columns.data_type 
		FROM information_schema.columns 
		WHERE table_schema = 'public' AND columns.table_name = $1
	`, tableName)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf(" Warning: Failed to close rows: %v", err)
		}
	}(rows)

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
