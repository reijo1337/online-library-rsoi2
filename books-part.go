package main

import (
	"context"
	"os"

	"github.com/reijo1337/online-library-rsoi2/books-service/protocol"
	"google.golang.org/grpc"
)

type BooksPart struct {
	conn  *grpc.ClientConn
	books protocol.BooksClient
}

func NewBooksPart() (*BooksPart, error) {
	addr := os.Getenv("BOOKSADDR")
	if addr == "" {
		addr = "0.0.0.0"
	}

	grpcConn, err := grpc.Dial(
		addr+":8081",
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, err
	}

	books := protocol.NewBooksClient(grpcConn)
	return &BooksPart{
		conn:  grpcConn,
		books: books,
	}, nil
}

func (bp *BooksPart) addNewBook(book Book) int32 {
	ctx := context.Background()

	insertBookRequest := &protocol.BookInsert{
		BookName:   book.Name,
		AuthorName: book.Author.Name,
	}

	id, err := bp.books.AddBook(ctx, insertBookRequest)

	if err != nil {
		panic(err)
	}

	return id.ID
}

func (bp *BooksPart) getBookByID(ID int32) (*Book, error) {
	ctx := context.Background()
	bookID := &protocol.SomeID{ID: ID}

	book, err := bp.books.BookByID(ctx, bookID)
	if err != nil {
		return nil, err
	}

	return &Book{
		ID:   book.GetID(),
		Name: book.GetName(),
		Author: &Writer{
			ID:   book.GetAuthor().GetID(),
			Name: book.GetAuthor().GetName(),
		},
		Free: book.GetFree(),
	}, nil
}

func (bp *BooksPart) changeBookStatusByID(ID int32, status bool) error {
	ctx := context.Background()
	req := &protocol.ChangeStatus{
		BookID:    ID,
		NewStatus: status,
	}

	_, err := bp.books.ChangeBookStatusByID(ctx, req)
	return err
}
