package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	USER     = "rsoi"
	PASSWORD = "password"
	DB_NAME  = "arrears"
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
		CREATE TABLE IF NOT EXISTS arrears (
			id SERIAL NOT NULL PRIMARY KEY,
			reader_id INTEGER NOT NULL,
			book_id INTEGER NOT NULL,
			start_date varchar(8) NOT NULL,
			end_date varchar(8) NOT NULL
		)`); err != nil {
		return err
	}

	return nil
}

func (db *Database) GetArrearsPaggin(userID int32, size int32, page int32) ([]*Arrear, error) {
	resultArrears := make([]*Arrear, 0)
	row, err := db.Query("SELECT * FROM arrears WHERE reader_id = $1 LIMIT $2 OFFSET $3", userID, size, (page-1)*size)
	if err != nil {
		return nil, err
	}
	currentArrear := &Arrear{}
	for row.Next() {
		row.Scan(&currentArrear.ID, &currentArrear.readerID, &currentArrear.bookID, &currentArrear.start, &currentArrear.end)
		resultArrears = append(resultArrears, currentArrear)
	}

	return resultArrears, nil
}
