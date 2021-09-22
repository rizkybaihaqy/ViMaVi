package config

import (
	"database/sql"
	"fmt"
)

// Create postgres db connection using .env variable
func CreateConnection(pgUrl string) *sql.DB {
	db, err := sql.Open("postgres", pgUrl)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB connection successfully!")
	return db
}
