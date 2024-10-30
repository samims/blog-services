package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"auth-service/constants"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func Connect() (*sql.DB, error) {
	// Define the connection string

	dbHost := os.Getenv(constants.PostgresHost)
	dbPort := os.Getenv(constants.PostgresPort)
	dbUser := os.Getenv(constants.PostgresUser)
	dbPass := os.Getenv(constants.PostgresPasswd)
	dbName := os.Getenv(constants.PostgresDBName)
	dbSSLMode := os.Getenv(constants.DatabaseSSLMode)

	pgInfo := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUser,
		dbPass,
		dbName,
		dbSSLMode,
	)

	log.Println("Connecting to database auth service", pgInfo)

	// Open a connection to the db
	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		log.Fatalf("error opening db %v", err)
		return nil, err
	}

	// check if the connection is valid
	if err = db.Ping(); err != nil {
		log.Fatalf("error conecting to the databse: %v", err)
		return nil, err
	}
	log.Println("Database connection established successfully")
	return db, nil

}
