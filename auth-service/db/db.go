package db

import (
	"auth-service/constants"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func Connect() (*sql.DB, error) {
	// Define the connection string
	//connStr := "postgresql://postgres@auth-service:5432/auth-service?sslmode=false"
	// read from env

	dbHost := os.Getenv(constants.PostgresHost)
	dbPort := os.Getenv(constants.PostgresPort)
	dbUser := os.Getenv(constants.PostgresUser)
	dbPass := os.Getenv(constants.PostgresPasswd)
	dbName := os.Getenv(constants.PostgresDBName)
	//dbSSLMode := os.Getenv(constants.DatabaseSSLMode)

	//connStr := fmt.Sprintf(
	//	"%s:%s@%s:%s/%s?sslmode=%s",
	//	dbUser,
	//	dbPass,
	//	dbHost,
	//	dbPort,
	//	dbName,
	//	dbSSLMode,
	//)
	pgInfo := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPass,
		dbName)

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
