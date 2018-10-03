package main

import (
	"context"

	"github.com/reijo1337/online-library-rsoi2/books-service/protocol"
)

type BooksServer struct {
	db *Database
}

// Server возвращает новый объект BookServer, который представляет определения для grpc
func Server() (*BooksServer, error) {
	db, err := SetUpDatabase()
	if err != nil {
		return nil, err
	}
	return &BooksServer{db: db}, nil
}

// Authors возвращает отправляет список авторов, чьи книги есть в библиотеке
func (s *BooksServer) Authors(in *protocol.Nothing, p protocol.Books_AuthorsServer) error {
	writers, err := s.writersList()
	if err != nil {
		return err
	}
	for _, writer := range writers {
		if err := p.Send(writer); err != nil {
			return err
		}
	}
	return nil
}

// BookByAuthorAndName возвращает книгу с определенным названием, определенного автора
func (s *BooksServer) BookByAuthorAndName(ctx context.Context, req *protocol.WriterBookName) (*protocol.Book, error) {
	book, err := s.db.getBookByNameAndAuthor(req.GetName(), req.GetWriter())
	if err != nil {
		return nil, err
	}
	writer := &protocol.Writer{ID: book.Author.ID, Name: book.Author.Name}
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

// Добавление новой книги в библиотеку
func (s *BooksServer) AddBook(ctx context.Context, bookInfo *protocol.BookInsert) (*protocol.SomeID, error) {
	newBookID, err := s.db.insertNewBook(bookInfo.BookName, bookInfo.AuthorName)

	if err != nil {
		return nil, err
	}

	return &protocol.SomeID{ID: newBookID.ID}, nil
}
