package main

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	USER     = "rsoi"
	PASSWORD = "password"
	DB_NAME  = "auth"
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
	db.SetMaxOpenConns(100)

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
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL NOT NULL PRIMARY KEY,
			login VARCHAR(20) NOT NULL,
			passhash VARCHAR(64) NOT NULL
		)`); err != nil {
		return err
	}

	return nil
}

func setUpStartData(db *Database) error {
	users := []User{
		User{
			Login:    "login1",
			Password: "password1",
		},
		User{
			Login:    "login2",
			Password: "password2",
		},
	}

	for _, d := range users {
		if _, err := db.InsertNewUser(d.Login, d.Password); err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) InsertNewUser(login string, password string) (*User, error) {
	log.Println("DB: Inserting new user ", login)
	passHash := sha256.New()
	passHash.Write([]byte(password))
	pass := passHash.Sum(nil)
	var ID int32
	row := db.QueryRow("INSERT INTO auth (login, passhash) VALUES ($1, $2,) RETURNING id",
		login, string(pass))
	if err := row.Scan(&ID); err != nil {
		return nil, err
	}
	log.Println("DB: user inserted succesfully")
	return &User{
		ID:       ID,
		Login:    login,
		Password: password,
	}, nil
}

func (db *Database) IsAuthorized(user *User) bool {
	log.Println("DB: Check user is in DB")
	var (
		passhash string
	)

	err := db.QueryRow("SELECT passhash FROM auth WHERE login = $1", user.Login).Scan(
		&passhash)
	if err != nil {
		return false
	}
	passHash := sha256.New()
	passHash.Write([]byte(user.Password))
	return bytes.Equal(passHash.Sum(nil), []byte(passhash))
}
