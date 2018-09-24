package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	USER     = "rsoi"
	PASSWORD = "password"
	DB_NAME  = "books"
)

type Database struct {
	*sql.DB
}

func SetUpDatabase() (*Database, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", USER, PASSWORD, DB_NAME))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	if err := createSchema(db); err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil

}

func createSchema(db *sql.DB) error {
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS writers (
			id SERIAL NOT NULL PRIMARY KEY,
			name VARCHAR(50) NOT NULL UNIQUE
		)`); err != nil {
		return err
	}

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id SERIAL NOT NULL PRIMARY KEY,
			author INT NOT NULL REFERENCES writers (id),
			name VARCHAR(50) NOT NULL UNIQUE
		)
	`); err != nil {
		return err
	}

	return nil
}

func setUpStartData(db *sql.DB) error {
	writer1 := &Writer{
		Name: "Edgar Allan Poe",
	}
	writer2 := &Writer{
		Name: "Howard Phillips Lovecraft",
	}
	writer3 := &Writer{
		Name: "Fyodor Dostoevsky",
	}

	books := []Book{
		Book{
			Name:   "The Black Cat",
			Author: writer1,
		},
		Book{
			Name:   "Morella",
			Author: writer1,
		},
		Book{
			Name:   "To Helen",
			Author: writer1,
		},
		Book{
			Name:   "Dagon",
			Author: writer2,
		},
		Book{
			Name:   "Memory",
			Author: writer2,
		},
		Book{
			Name:   "The Shadow Out of Time",
			Author: writer2,
		},
		Book{
			Name:   "Poor Folk",
			Author: writer3,
		},
		Book{
			Name:   "The Idiot",
			Author: writer3,
		},
		Book{
			Name:   "The Brothers Karamazov",
			Author: writer3,
		},
	}
}
