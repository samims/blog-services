package db

import (
	"database/sql"
	"log"
)

func Connect() (*sql.DB, error) {
	// Define the connection string
	connStr := "postgresql://postgres@auth-service:5432/auth-service?sslmode=false"

	// Open a connection to the db
	db, err := sql.Open("postgres", connStr)
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
