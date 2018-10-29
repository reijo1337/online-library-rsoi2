package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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
			name VARCHAR(50) NOT NULL UNIQUE,
			free BOOLEAN DEFAULT TRUE
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
	log.Println("DB: Inserting new writer named", name)
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

	log.Println("DB: writer inserted succesfully")
	return &Writer{ID: ID, Name: name}, nil
}

func (db *Database) insertBook(name string, writer *Writer) (*Book, error) {
	log.Println("DB: Inserting new book named", name, "written by", writer.Name)
	rows, err := db.Query("SELECT id FROM books WHERE name = $1 AND author = $2", name, writer.ID)

	if err != nil {
		return nil, err
	}

	var ID int32
	for rows.Next() {
		rows.Scan(&ID)
	}

	if ID > 0 {
		return &Book{ID: ID, Name: name, Author: writer, Free: true}, nil
	}

	row := db.QueryRow("INSERT INTO books (name, author) VALUES ($1, $2) RETURNING id", name, writer.ID)

	if err := row.Scan(&ID); err != nil {
		return nil, err
	}

	log.Println("DB: book inserted succesfully")
	return &Book{ID: ID, Name: name, Author: writer}, nil
}

func (db *Database) getAllAuthors() ([]*Writer, error) {
	log.Println("DB: Getting all writers")
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

	log.Println("DB: writers received succesfully")
	return resultWriters, nil
}

func (db *Database) getBookByNameAndAuthor(name string, author string) (*Book, error) {
	log.Println("DB: Getting book named", name, "written by", author)
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

	rows, err = db.Query("SELECT id, free FROM books WHERE name = $1 AND author = $2", name, ID)

	if err != nil {
		return nil, err
	}

	ID = 0
	var free bool
	for rows.Next() {
		rows.Scan(&ID, free)
	}

	if ID > 0 {
		log.Println("DB: book received succesfully")
		return &Book{
				ID:     ID,
				Name:   name,
				Author: returnWriter,
				Free:   free},
			nil
	}

	return nil, errors.New("There is no book with name " + name)
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
	log.Println("DB: Getting book with ID", ID)
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
	log.Println("DB: book received succesfully")
	return &Book{
		ID:   ID,
		Name: name,
		Author: &Writer{
			ID:   authorID,
			Name: authorName,
		},
	}, nil
}

func (db *Database) changeStatusBookByID(ID int32, status bool) (bool, error) {
	log.Println("DB: Changing 'free' status of book with ID", ID, "to", status)
	var free bool
	err := db.QueryRow("SELECT free FROM books where id = $1", ID).Scan(&free)
	if err != nil {
		return false, err
	}

	if free == status {
		return false, errors.New("This book already in your state")
	}

	_, err = db.Query("UPDATE books SET free = $2 WHERE id = $1", ID, status)
	if err != nil {
		return false, err
	}
	log.Println("DB: Book status changed successfully")
	return true, nil
}
