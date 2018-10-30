package main

import (
	"context"
	"log"
	"os"

	"github.com/reijo1337/online-library-rsoi2/books-service/protocol"
	"google.golang.org/grpc"
)

type BooksPart struct {
	conn  *grpc.ClientConn
	books protocol.BooksClient
}

func NewBooksPart() (*BooksPart, error) {
	log.Println("Books Client: Connecting to books service...")
	addr := os.Getenv("BOOKSADDR")
	if addr == "" {
		addr = "0.0.0.0"
	}

	log.Println("Books Client: books service addres:", addr+":8081")
	grpcConn, err := grpc.Dial(
		addr+":8081",
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Println("Books Client: Can't connect to remote service")
		return nil, err
	}

	books := protocol.NewBooksClient(grpcConn)
	log.Println("Books Client: success!")
	return &BooksPart{
		conn:  grpcConn,
		books: books,
	}, nil
}

func (bp *BooksPart) addNewBook(book Book) (int32, error) {
	log.Println("Books Client: adding new book named", book.Name, "by", book.Author.Name)
	ctx := context.Background()

	insertBookRequest := &protocol.BookInsert{
		BookName:   book.Name,
		AuthorName: book.Author.Name,
	}

	id, err := bp.books.AddBook(ctx, insertBookRequest)

	if err != nil {
		log.Println("Books Client: Can't add new book")
		return 0, err
	}

	log.Println("Books Client: Book added succesfully")
	return id.ID, nil
}

func (bp *BooksPart) getBookByID(ID int32) (*Book, error) {
	log.Println("Books Client: Getting book with ID", ID)
	ctx := context.Background()
	bookID := &protocol.SomeID{ID: ID}

	book, err := bp.books.BookByID(ctx, bookID)
	if err != nil {
		log.Println("Books Client: Can't get book")
		return nil, err
	}

	log.Println("Books Client: Book getted succesfully")
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
	log.Println("Books Client: Changing book status to", status, ". Book ID:", ID)
	ctx := context.Background()
	req := &protocol.ChangeStatus{
		BookID:    ID,
		NewStatus: status,
	}

	_, err := bp.books.ChangeBookStatusByID(ctx, req)
	if err != nil {
		log.Println("Books Client: Can't change book status")
	} else {
		log.Println("Books Client: Status changed succesfully")
	}
	return err
}
