package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	USER     = "postgres"
	PASSWORD = "password"
	DB_NAME  = "postgres"
)

type Database struct {
	*sql.DB
}

func SetUpDatabase() (*Database, error) {
	log.Println("DB: Connecting to", DB_NAME, "database")
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=db", USER, PASSWORD, DB_NAME))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(100)

	log.Println("Creating schema")
	if err := createSchema(db); err != nil {
		return nil, err
	}

	ddb := &Database{DB: db}

	// log.Println("DB: Setting up start data")
	// if err := setUpStartData(ddb); err != nil {
	// 	return nil, err
	// }

	log.Println("DB: succesful setup")
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

func setUpStartData(db *Database) error {
	data := []Arrear{
		Arrear{
			readerID: 1,
			bookID:   2,
			start:    "20181024",
			end:      "20181124",
		},
		Arrear{
			readerID: 2,
			bookID:   3,
			start:    "20181024",
			end:      "20181124",
		},
		Arrear{
			readerID: 3,
			bookID:   4,
			start:    "20181024",
			end:      "20181124",
		},
	}
	for _, d := range data {
		if _, err := db.InsertNewArrear(d.readerID, d.bookID, d.start, d.end); err != nil {
			return err
		}
	}
	return nil
}

func (db *Database) GetArrearsPaggin(userID int32, size int32, page int32) ([]*Arrear, error) {
	log.Println("DB: Getting", page, "page of arrears with reader id", userID, "with page size", size)
	resultArrears := make([]*Arrear, 0)
	row, err := db.Query("SELECT * FROM arrears WHERE reader_id = $1 LIMIT $2 OFFSET $3", userID, size, (page-1)*size)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		currentArrear := &Arrear{}
		row.Scan(&currentArrear.ID, &currentArrear.readerID, &currentArrear.bookID, &currentArrear.start, &currentArrear.end)
		resultArrears = append(resultArrears, currentArrear)
	}

	log.Println("DB: arrears received succesfully")
	return resultArrears, nil
}

func (db *Database) InsertNewArrear(readerID int32, bookID int32, startTime string, endTime string) (*Arrear, error) {
	log.Println("DB: Inserting new arrear with", readerID, "reader id and ", bookID, "book id for period [", startTime, ";", endTime, "]")
	var ID int32
	row := db.QueryRow("INSERT INTO arrears (reader_id, book_id, start_date, end_date) VALUES ($1, $2, $3, $4) RETURNING id",
		readerID, bookID, startTime, endTime)
	if err := row.Scan(&ID); err != nil {
		return nil, err
	}
	log.Println("DB: arrear inserted succesfully")
	return &Arrear{
		ID:       ID,
		readerID: readerID,
		bookID:   bookID,
		start:    startTime,
		end:      endTime,
	}, nil
}

func (db *Database) InsertFullArrear(ID int32, readerID int32, bookID int32, startTime string, endTime string) (*Arrear, error) {
	log.Println("DB: Inserting new arrear with", readerID, "reader id and ", bookID, "book id for period [", startTime, ";", endTime, "]")

	rows, err := db.Query("SELECT * FROM arrears WHERE id = $1", ID)

	if err != nil {
		log.Println("DB: Can't check arrear for already exist:", err.Error())
		return nil, err
	}

	ar := &Arrear{ID: int32(-1)}
	for rows.Next() {
		rows.Scan(&ar.ID, &ar.readerID, &ar.bookID, &ar.start, &ar.end)
	}

	if ID > 0 {
		log.Println("DB: There is already arrear with this ID")
		return ar, nil
	}
	db.QueryRow("INSERT INTO arrears (id, reader_id, book_id, start_date, end_date) VALUES ($1, $2, $3, $4)",
		ID, readerID, bookID, startTime, endTime)

	log.Println("DB: arrear inserted succesfully")
	return &Arrear{
		ID:       ID,
		readerID: readerID,
		bookID:   bookID,
		start:    startTime,
		end:      endTime,
	}, nil
}

func (db *Database) GetArrearByID(ID int32) (*Arrear, error) {
	log.Println("DB: Getting arrear with ID", ID)
	var (
		readerID int32
		bookID   int32
		start    string
		end      string
	)

	err := db.QueryRow("SELECT reader_id, book_id, start_date, end_date FROM arrears WHERE id = $1", ID).Scan(
		&readerID, &bookID, &start, &end)
	if err != nil {
		return nil, err
	}
	log.Println("DB: arrear received succesfully")
	return &Arrear{
		ID:       ID,
		readerID: readerID,
		bookID:   bookID,
		start:    start,
		end:      end,
	}, nil
}

func (db *Database) CloseArrayByID(ID int32) error {
	log.Println("DB: deleting arrear with ID", ID)
	_, err := db.Query("DELETE FROM arrears WHERE id = $1", ID)
	return err
}
