package clients

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc/metadata"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/reijo1337/online-library-rsoi2/secrets"

	"github.com/reijo1337/online-library-rsoi2/books-service/protocol"
	"google.golang.org/grpc"
)

type BooksPartInterface interface {
	AddNewBook(book Book) (int32, error)
	GetBookByID(ID int32) (*Book, error)
	ChangeBookStatusByID(ID int32, status bool) error
	GetFreeBooks() ([]Book, error)
}

type BooksPart struct {
	conn  *grpc.ClientConn
	books protocol.BooksClient
	token string
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

	log.Println("Books Client: Getting access token")
	token, err := Auth(books)
	if err != nil {
		log.Println("Books Client: Can't authorize")
		return nil, err
	}

	bp := &BooksPart{
		conn:  grpcConn,
		books: books,
		token: token,
	}

	go WaithAndRefresh(bp)

	return bp, nil
}

func (bp *BooksPart) Context() context.Context {
	header := metadata.New(map[string]string{"authorization": bp.token})
	return metadata.NewOutgoingContext(context.Background(), header)
}

func (bp *BooksPart) AddNewBook(book Book) (int32, error) {
	log.Println("Books Client: adding new book named", book.Name, "by", book.Author.Name)
	ctx := bp.Context()

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

func (bp *BooksPart) GetBookByID(ID int32) (*Book, error) {
	log.Println("Books Client: Getting book with ID", ID)
	ctx := bp.Context()
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

func (bp *BooksPart) ChangeBookStatusByID(ID int32, status bool) error {
	log.Println("Books Client: Changing book status to", status, ". Book ID:", ID)
	ctx := bp.Context()
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

func (bp *BooksPart) GetFreeBooks() ([]Book, error) {
	log.Println("Books Client: Getting free books")
	ctx := bp.Context()
	in := &protocol.NothingBooks{}
	booksServ, err := bp.books.FreeBooks(ctx, in)
	if err != nil {
		log.Println("Books Client: Can't recieve book")
		return nil, err
	}
	var books []Book

	for {
		recvBook, err := booksServ.Recv()
		if err == io.EOF {
			log.Println("Arrear Client: All arrears recieved successfully")
			return books, nil
		} else if err != nil {
			log.Println("Arrear Client: Can't receive arrear")
			return nil, err
		}
		books = append(books,
			Book{
				ID:   recvBook.GetID(),
				Name: recvBook.GetName(),
				Free: recvBook.GetFree(),
				Author: &Writer{
					ID:   recvBook.Author.GetID(),
					Name: recvBook.Author.GetName(),
				},
			})
	}
}

func Auth(books protocol.BooksClient) (string, error) {
	log.Println("Books Client: Getting authorization")

	ctx := context.Background()
	request := &protocol.AuthRequest{
		AppKey:    secrets.AppKey,
		AppSecret: secrets.BooksSecret,
	}
	log.Println("Books Client: Sending request for authorization")
	token, err := books.Auth(ctx, request)
	if err != nil {
		log.Println("Books Client: Can't authorize: ", err.Error())
		return "", err
	}

	return token.GetString_(), nil
}

func WaithAndRefresh(bp *BooksPart) {
	tokenString := bp.token
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secrets.BooksSecret), nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)
	endTime := int64(claims["exp"].(float64))
	startTime := time.Now().Unix()
	seconsd := int(endTime - startTime)
	time.Sleep(time.Duration(seconsd) * time.Second)

	newToken, _ := Auth(bp.books)
	bp.token = newToken
	go WaithAndRefresh(bp)
}
