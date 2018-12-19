package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/metadata"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/reijo1337/online-library-rsoi2/books-service/protocol"
)

type BooksServer struct {
	db *Database
}

// Server возвращает новый объект BookServer, который представляет определения для grpc
func Server() (*BooksServer, error) {
	log.Println("Set up book service...")
	db, err := SetUpDatabase()
	if err != nil {
		return nil, err
	}
	return &BooksServer{db: db}, nil
}

func isAuthorized(ctx context.Context) bool {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ok
	}

	tokenString := headers["authorization"][0]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(Secret), nil
	})

	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	} else {
		return false
	}
}

// Authors возвращает отправляет список авторов, чьи книги есть в библиотеке
func (s *BooksServer) Authors(in *protocol.NothingBooks, p protocol.Books_AuthorsServer) error {
	if !isAuthorized(p.Context()) {
		return errors.New("Unauthorized")
	}
	log.Println("Server: New request for writers list")
	writers, err := s.writersList()
	if err != nil {
		log.Println("Server: Can't process this request:", err.Error())
		return err
	}
	for _, writer := range writers {
		if err := p.Send(writer); err != nil {
			log.Println("Server: Can't send writeer:", err.Error())
			return err
		}
	}
	log.Println("Server: Request processed successfully")
	return nil
}

// BookByAuthorAndName возвращает книгу с определенным названием, определенного автора
func (s *BooksServer) BookByAuthorAndName(ctx context.Context, req *protocol.WriterBookName) (*protocol.Book, error) {
	if !isAuthorized(ctx) {
		return nil, errors.New("Unauthorized")
	}
	log.Println("Server: New request for book named", req.GetName(), "written by", req.GetWriter())
	book, err := s.db.getBookByNameAndAuthor(req.GetName(), req.GetWriter())
	if err != nil {
		log.Println("Server: Can't process this request:", err.Error())
		return nil, err
	}
	writer := &protocol.Writer{ID: book.Author.ID, Name: book.Author.Name}
	log.Println("Server: Request processed successfully")
	return &protocol.Book{ID: book.ID, Name: book.Name, Author: writer}, nil
}

func (s *BooksServer) writersList() ([]*protocol.Writer, error) {
	writers, err := s.db.getAllAuthors()
	if err != nil {
		return nil, err
	}
	ret := make([]*protocol.Writer, len(writers))
	for _, writer := range writers {
		ret = append(ret, &protocol.Writer{ID: writer.ID, Name: writer.Name})
	}

	return ret, nil
}

func (s *BooksServer) freeBooksList() ([]*protocol.Book, error) {
	books, err := s.db.getFreeBooks()
	if err != nil {
		return nil, err
	}
	ret := make([]*protocol.Book, len(books))
	for _, book := range books {
		writer := &protocol.Writer{ID: book.Author.ID, Name: book.Author.Name}
		ret = append(ret, &protocol.Book{ID: book.ID, Name: book.Name, Author: writer, Free: book.Free})
	}

	return ret, nil
}

// AddBook добавление новой книги в библиотеку
func (s *BooksServer) AddBook(ctx context.Context, bookInfo *protocol.BookInsert) (*protocol.SomeID, error) {
	if !isAuthorized(ctx) {
		return nil, errors.New("Unauthorized")
	}
	log.Println("Server: New request for inserting book named", bookInfo.BookName, "written by", bookInfo.AuthorName)
	newBookID, err := s.db.insertNewBook(bookInfo.BookName, bookInfo.AuthorName)

	if err != nil {
		log.Println("Server: Can't process request:", err.Error())
		return nil, err
	}

	log.Println("Server: Request processed successfully")
	return &protocol.SomeID{ID: newBookID.ID}, nil
}

// BookByID возвращает книгу по ID
func (s *BooksServer) BookByID(ctx context.Context, req *protocol.SomeID) (*protocol.Book, error) {
	if !isAuthorized(ctx) {
		return nil, errors.New("Unauthorized")
	}
	log.Println("Server: New request for book with ID", req.GetID())
	book, err := s.db.getBookByID(req.GetID())
	if err != nil {
		log.Println("Server: Can't process request:", err.Error())
		return nil, err
	}
	writer := &protocol.Writer{ID: book.Author.ID, Name: book.Author.Name}
	log.Println("Server: Request processed successfully")
	return &protocol.Book{ID: book.ID, Name: book.Name, Author: writer}, nil
}

// ChangeBookStatusByID изменяет статус книги "занята или нет"
func (s *BooksServer) ChangeBookStatusByID(ctx context.Context, in *protocol.ChangeStatus) (*protocol.NothingBooks, error) {
	if !isAuthorized(ctx) {
		return nil, errors.New("Unauthorized")
	}
	log.Println("Server: New request for changing 'free' book status to", in.GetNewStatus(), ", ID", in.GetBookID())
	changed, err := s.db.changeStatusBookByID(in.GetBookID(), in.GetNewStatus())
	if err != nil {
		log.Println("Server: Can't process request:", err.Error())
	} else {
		log.Println("Server: Request processed successfully")
	}
	return &protocol.NothingBooks{Dummy: changed}, err
}

// FreeBooks возвращает список свободных книг, т.е. книг, которые находядтся непостредственно в библиотеке
func (s *BooksServer) FreeBooks(in *protocol.NothingBooks, p protocol.Books_FreeBooksServer) error {
	if !isAuthorized(p.Context()) {
		return errors.New("Unauthorized")
	}
	log.Println("Server: New request for free book list")
	books, err := s.freeBooksList()
	if err != nil {
		log.Println("Server: Can't process this request:", err.Error())
		return err
	}
	for _, book := range books {
		if book == nil {
			continue
		}
		if err := p.Send(book); err != nil {
			log.Println("Server: Can't send book:", err.Error())
			return err
		}
	}
	log.Println("Server: Request processed successfully")
	return nil
}

func (s *BooksServer) Auth(ctx context.Context, in *protocol.AuthRequest) (*protocol.SomeString, error) {
	log.Println("Server: New authorization")
	if ContainsKey(in.AppKey) && (in.AppSecret == Secret) {
		token, err := genToken(in.AppKey)
		return &protocol.SomeString{String_: token}, err
	}
	return nil, errors.New("Unauthorized")
}

func genToken(login string) (string, error) {
	log.Println("Server: Generating token")
	hmacSampleSecret := []byte(Secret)
	AccessTokenExp := time.Now().Add(time.Minute * 30).Unix()
	log.Println("Server: Gen access token")
	accesToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "BookService",
		"exp": AccessTokenExp,
		"aud": login,
	})

	log.Println("Server: Signing access token", accesToken, hmacSampleSecret)
	accessTokenString, err := accesToken.SignedString(hmacSampleSecret)
	if err != nil {
		log.Println("Server: Can't authorize: ", err.Error())
		return "", err
	}

	return accessTokenString, nil
}
