package main

import (
	"context"
	"log"

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

// Authors возвращает отправляет список авторов, чьи книги есть в библиотеке
func (s *BooksServer) Authors(in *protocol.NothingBooks, p protocol.Books_AuthorsServer) error {
	log.Println("Server: New request for writers list")
	writers, err := s.writersList()
	if err != nil {
		log.Fatalln("Server: Can't process this request:", err.Error())
		return err
	}
	for _, writer := range writers {
		if err := p.Send(writer); err != nil {
			log.Fatalln("Server: Can't send writeer:", err.Error())
			return err
		}
	}
	log.Println("Server: Request processed successfully")
	return nil
}

// BookByAuthorAndName возвращает книгу с определенным названием, определенного автора
func (s *BooksServer) BookByAuthorAndName(ctx context.Context, req *protocol.WriterBookName) (*protocol.Book, error) {
	log.Println("Server: New request for book named", req.GetName(), "written by", req.GetWriter())
	book, err := s.db.getBookByNameAndAuthor(req.GetName(), req.GetWriter())
	if err != nil {
		log.Fatalln("Server: Can't process this request:", err.Error())
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

// AddBook добавление новой книги в библиотеку
func (s *BooksServer) AddBook(ctx context.Context, bookInfo *protocol.BookInsert) (*protocol.SomeID, error) {
	log.Println("Server: New request for inserting book named", bookInfo.BookName, "written by", bookInfo.AuthorName)
	newBookID, err := s.db.insertNewBook(bookInfo.BookName, bookInfo.AuthorName)

	if err != nil {
		log.Fatalln("Server: Can't process request:", err.Error())
		return nil, err
	}

	log.Println("Server: Request processed successfully")
	return &protocol.SomeID{ID: newBookID.ID}, nil
}

// BookByID возвращает книгу по ID
func (s *BooksServer) BookByID(ctx context.Context, req *protocol.SomeID) (*protocol.Book, error) {
	log.Println("Server: New request for book with ID", req.GetID())
	book, err := s.db.getBookByID(req.GetID())
	if err != nil {
		log.Fatalln("Server: Can't process request:", err.Error())
		return nil, err
	}
	writer := &protocol.Writer{ID: book.Author.ID, Name: book.Author.Name}
	log.Println("Server: Request processed successfully")
	return &protocol.Book{ID: book.ID, Name: book.Name, Author: writer}, nil
}

// ChangeBookStatusByID изменяет статус книги "занята или нет"
func (s *BooksServer) ChangeBookStatusByID(ctx context.Context, in *protocol.ChangeStatus) (*protocol.NothingBooks, error) {
	log.Println("Server: New request for changing 'free' book status to", in.GetNewStatus(), ", ID", in.GetBookID())
	changed, err := s.db.changeStatusBookByID(in.GetBookID(), in.GetNewStatus())
	if err != nil {
		log.Fatalln("Server: Can't process request:", err.Error())
	} else {
		log.Println("Server: Request processed successfully")
	}
	return &protocol.NothingBooks{Dummy: changed}, err
}
