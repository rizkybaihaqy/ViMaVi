package config

import (
	"database/sql"
	"fmt"
)

// Create postgres db connection using .env variable
func CreateConnection(pgUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", pgUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("DB connection successfully!")
	return db, nil
}
