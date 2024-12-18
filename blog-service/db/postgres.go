package db

import (
	"blog-service/logger"
	"context"
	"database/sql"
	"fmt"

	"blog-service/config"

	_ "github.com/lib/pq"
)

// PostgresConnector interface defines methods that any database connector must implement
type PostgresConnector interface {
	Connect(ctx context.Context) (*sql.DB, error)
	Disconnect() error
	IsConnected(ctx context.Context) bool
	GetConfig() config.PostgresConfig
	GetDB() *sql.DB
}

// postgresConnector provides common functionalities for PostgreSQL connectors
type postgresConnector struct {
	pgConf config.PostgresConfig
	db     *sql.DB
	log    *logger.AppLogger
}

// NewPostgresConnector creates a new instance of postgresConnector
func NewPostgresConnector(cfg config.PostgresConfig, log *logger.AppLogger) PostgresConnector {
	return &postgresConnector{
		pgConf: cfg,
		log:    log,
	}
}

// Connect establishes a connection to the PostgreSQL database
func (pc *postgresConnector) Connect(ctx context.Context) (*sql.DB, error) {
	// Create connection string and establish a new connection
	connStr := pc.pgConf.ConnectionURL()

	pc.log.Debugf("Connecting to postgres at %s", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database %w", err)
	}
	pc.log.Debugf("Successfully opened db instance postgres at %s", connStr)

	if err = db.PingContext(ctx); err != nil {
		pc.log.Warnf("failed to ping database connection with connectionString %s", connStr)
		closeErr := db.Close()
		if closeErr != nil {
			pc.log.WithContext(ctx).Error("failed to close database connection")
		}
		return nil, fmt.Errorf("failed to ping database %w", err)
	}

	pc.db = db
	pc.log.Println("successfully connected to postgres...")
	return db, nil
}

// Disconnect closes the database connection
func (pc *postgresConnector) Disconnect() error {
	if pc.GetDB() != nil {
		closeErr := pc.GetDB().Close()
		if closeErr != nil {
			pc.log.Warn(context.Background(), "failed to close database connection")
			return closeErr
		}
		pc.log.Info(context.Background(), "successfully closed database connection")
	}
	return nil
}

// IsConnected checks if the database connection is active
func (pc *postgresConnector) IsConnected(ctx context.Context) bool {
	if pc.db == nil {
		return false
	}
	err := pc.db.PingContext(ctx)
	if err != nil {
		pc.log.Warnf("failed to ping database connection %v", err)
		return false
	}
	return true
}

// GetConfig returns the PostgresConfig of the connector
func (pc *postgresConnector) GetConfig() config.PostgresConfig {
	return pc.pgConf
}

// GetDB returns the current database connection
func (pc *postgresConnector) GetDB() *sql.DB {
	if pc.db == nil {
		return nil
	}
	return pc.db
}
