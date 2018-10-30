package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	USER     = "rsoi"
	PASSWORD = "password"
	DB_NAME  = "readers"
)

type Database struct {
	*sql.DB
}

func SetUpDatabase() (*Database, error) {
	log.Println("DB: Connecting to", DB_NAME, "database")
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", USER, PASSWORD, DB_NAME))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	log.Println("Creating schema")
	if err := createSchema(db); err != nil {
		return nil, err
	}

	ddb := &Database{DB: db}

	log.Println("DB: Setting up start data")
	if err := setUpStartData(ddb); err != nil {
		return nil, err
	}

	log.Println("DB: succesful setup")
	return ddb, nil

}

func setUpStartData(db *Database) error {
	names := []string{
		"Jhon Snow",
		"Vladimir Putin",
		"Ivan Ivanov",
	}

	for _, name := range names {
		if _, err := db.addReader(name); err != nil {
			return err
		}
	}
	return nil
}

func createSchema(db *sql.DB) error {
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS readers (
			id SERIAL NOT NULL PRIMARY KEY,
			name VARCHAR(50) NOT NULL UNIQUE
		)`); err != nil {
		return err
	}

	return nil
}

func (db *Database) addReader(name string) (*Reader, error) {
	log.Println("DB: Adding new reader", name)
	rows, err := db.Query("SELECT id FROM readers WHERE name = $1", name)

	if err != nil {
		log.Println("DB: Can't check user for already exist:", err.Error())
		return nil, err
	}

	var ID int32
	for rows.Next() {
		rows.Scan(&ID)
	}

	if ID > 0 {
		log.Println("DB: There is already user with name", name)
		return &Reader{ID: ID, Name: name}, nil
	}

	row := db.QueryRow("INSERT INTO readers (name) VALUES ($1) RETURNING id", name)

	if err := row.Scan(&ID); err != nil {
		log.Println("DB: Can't insert user:", err.Error())
		return nil, err
	}

	return &Reader{ID: ID, Name: name}, nil
}

func (db *Database) getReadersList() ([]*Reader, error) {
	resultWriters := make([]*Reader, 0)
	rows, err := db.Query("SELECT * FROM positions ORDER BY time DESC")

	if err != nil {
		return nil, err
	}

	currentWriterInRows := &Reader{}
	for rows.Next() {
		rows.Scan(&currentWriterInRows.ID, &currentWriterInRows.Name)

		resultWriters = append(resultWriters, currentWriterInRows)
	}

	return resultWriters, nil
}

func (db *Database) getReaderByName(name string) (*Reader, error) {
	rows, err := db.Query("SELECT id FROM readers WHERE name = $1", name)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var ID int32
	for rows.Next() {
		rows.Scan(&ID)
	}

	if ID > 0 {
		return &Reader{ID: ID, Name: name}, nil
	} else {
		log.Println("No reader with name", name)
		return nil, errors.New("No reader with name " + name)
	}
}

func (db *Database) getReaderByID(ID int32) (*Reader, error) {
	rows, err := db.Query("SELECT name FROM readers WHERE id = $1", ID)

	if err != nil {
		return nil, err
	}

	var name string
	for rows.Next() {
		rows.Scan(&name)
	}

	if name != "" {
		return &Reader{ID: ID, Name: name}, nil
	} else {
		return nil, errors.New("No such reader")
	}
}
