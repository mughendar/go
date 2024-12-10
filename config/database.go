package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB initializes the database connection
func InitDB() (*sql.DB, error) {
	// Replace with your MySQL credentials
	dsn := "username:password@tcp(localhost:3306)/dbname"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Verify the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	fmt.Println("Database connected successfully!")
	return db, nil
}
