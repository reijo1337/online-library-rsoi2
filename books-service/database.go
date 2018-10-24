package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

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

	ddb := &Database{DB: db}

	if err := setUpStartData(ddb); err != nil {
		return nil, err
	}

	return ddb, nil

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

func setUpStartData(db *Database) error {
	writer1, err := db.insertWriter("Edgar Allan Poe")
	if err != nil {
		return err
	}
	writer2, err := db.insertWriter("Howard Phillips Lovecraft")
	if err != nil {
		return err
	}
	writer3, err := db.insertWriter("Fyodor Dostoevsky")
	if err != nil {
		return err
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

	for _, book := range books {
		if _, err := db.insertBook(book.Name, book.Author); err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) insertWriter(name string) (*Writer, error) {
	rows, err := db.Query("SELECT id FROM writers WHERE name = $1", name)

	if err != nil {
		return nil, err
	}

	var ID int32
	for rows.Next() {
		rows.Scan(&ID)
	}

	if ID > 0 {
		return &Writer{ID: ID, Name: name}, nil
	}

	row := db.QueryRow("INSERT INTO writers (name) VALUES ($1) RETURNING id", name)

	if err := row.Scan(&ID); err != nil {
		return nil, err
	}

	return &Writer{ID: ID, Name: name}, nil
}

func (db *Database) insertBook(name string, writer *Writer) (*Book, error) {
	rows, err := db.Query("SELECT id FROM books WHERE name = $1 AND author = $2", name, writer.ID)

	if err != nil {
		return nil, err
	}

	var ID int32
	for rows.Next() {
		rows.Scan(&ID)
	}

	if ID > 0 {
		return &Book{ID: ID, Name: name, Author: writer}, nil
	}

	row := db.QueryRow("INSERT INTO books (name, author) VALUES ($1, $2) RETURNING id", name, writer.ID)

	if err := row.Scan(&ID); err != nil {
		return nil, err
	}

	return &Book{ID: ID, Name: name, Author: writer}, nil
}

func (db *Database) getAllAuthors() ([]*Writer, error) {
	resultWriters := make([]*Writer, 0)
	rows, err := db.Query("SELECT * FROM positions ORDER BY time DESC")

	if err != nil {
		return nil, err
	}

	currentWriterInRows := &Writer{}
	for rows.Next() {
		rows.Scan(&currentWriterInRows.ID, &currentWriterInRows.Name)

		resultWriters = append(resultWriters, currentWriterInRows)
	}

	return resultWriters, nil
}

func (db *Database) getBookByNameAndAuthor(name string, author string) (*Book, error) {
	rows, err := db.Query("SELECT id FROM writers WHERE name = $1", author)

	if err != nil {
		return nil, err
	}

	var ID int32
	for rows.Next() {
		rows.Scan(&ID)
	}

	if ID == 0 {
		return nil, errors.New("There is no writer named " + author)
	}

	returnWriter := &Writer{ID: ID, Name: author}

	rows, err = db.Query("SELECT id FROM books WHERE name = $1 AND author = $2", name, ID)

	if err != nil {
		return nil, err
	}

	ID = 0
	for rows.Next() {
		rows.Scan(&ID)
	}

	returnBook := &Book{ID: ID, Name: name, Author: returnWriter}

	if ID > 0 {
		return returnBook, nil
	}

	return nil, nil
}

func (db *Database) insertNewBook(bookName string, authorName string) (*Book, error) {
	author, err := db.insertWriter(authorName)
	if err != nil {
		return nil, err
	}

	book, err := db.insertBook(bookName, author)
	return book, err
}

func (db *Database) getBookByID(ID int32) (*Book, error) {
	rows, err := db.Query("SELECT name, author FROM books where id = $1", ID)
	if err != nil {
		return nil, err
	}
	var (
		name     string
		authorID int32
	)
	for rows.Next() {
		rows.Scan(&name, &authorID)
	}
	if authorID == 0 {
		return nil, errors.New("There is no books with ID " + strconv.Itoa(int(ID)))
	}

	rows, err = db.Query("SELECT name FROM writers WHERE id = $1", authorID)

	if err != nil {
		return nil, err
	}

	var authorName string
	for rows.Next() {
		rows.Scan(&authorName)
	}
	return &Book{
		ID:   ID,
		Name: name,
		Author: &Writer{
			ID:   authorID,
			Name: authorName,
		},
	}, nil
}
