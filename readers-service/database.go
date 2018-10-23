package main

import (
	"database/sql"
	"errors"
	"fmt"

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
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", USER, PASSWORD, DB_NAME))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	if err := createSchema(db); err != nil {
		return nil, err
	}

	ddb := &Database{DB: db}

	//	if err := setUpStartData(ddb); err != nil {
	//		return nil, err
	//	}

	return ddb, nil

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
	rows, err := db.Query("SELECT id FROM readers WHERE name = $1", name)

	if err != nil {
		return nil, err
	}

	var ID int32
	for rows.Next() {
		rows.Scan(&ID)
	}

	if ID > 0 {
		return &Reader{ID: ID, Name: name}, nil
	}

	row := db.QueryRow("INSERT INTO readers (name) VALUES ($1) RETURNING id", name)

	if err := row.Scan(&ID); err != nil {
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
		return nil, err
	}

	var ID int32
	for rows.Next() {
		rows.Scan(&ID)
	}

	if ID > 0 {
		return &Reader{ID: ID, Name: name}, nil
	} else {
		return nil, errors.New("No such reader")
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
