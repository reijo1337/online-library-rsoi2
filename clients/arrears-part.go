package clients

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/reijo1337/online-library-rsoi2/arrears-service/protocol"
	"github.com/reijo1337/online-library-rsoi2/secrets"
	"google.golang.org/grpc"
)

type ArrearsPartInterface interface {
	GetArrearsPaging(userID int32, page int32, size int32) ([]Arrear, error)
	NewArrear(readerID int32, bookID int32) (*Arrear, error)
	GetArrearByID(ID int32) (*Arrear, error)
	CloseArrearByID(ID int32) error
}

type ArrearsPart struct {
	conn    *grpc.ClientConn
	arrears protocol.ArrearsClient
	token   string
}

func NewArrearsPart() (*ArrearsPart, error) {
	log.Println("Arrear Client: Connecting to arrear service...")
	addr := os.Getenv("ARREARSADDR")
	if addr == "" {
		addr = "0.0.0.0"
	}

	log.Println("Arrear Client: arrear service addres:", addr+":8083")
	grpcConn, err := grpc.Dial(
		addr+":8083",
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Println("Arrear Client: Can't connect to remote service")
		return nil, err
	}

	arrears := protocol.NewArrearsClient(grpcConn)
	log.Println("Arrear Client: success!")

	log.Println("Arrear Client: Getting access token")
	token, err := AuthA(arrears)
	if err != nil {
		log.Println("Books Client: Can't authorize")
		return nil, err
	}

	as := &ArrearsPart{
		conn:    grpcConn,
		arrears: arrears,
		token:   token,
	}

	go WaithAndRefreshA(as)

	return as, nil
}

func (ap *ArrearsPart) GetArrearsPaging(userID int32, page int32, size int32) ([]Arrear, error) {
	log.Println("Arrear Client: Getting arrears with pagging. User ID:", userID, ", page:", page, ", page size:", size)
	// header := metadata.New(map[string]string{"Authorization": secrets.AppKey + ":" + secrets.ArrearsSecret})
	// ctx := metadata.NewOutgoingContext(context.Background(), header)
	ctx := context.Background()
	in := &protocol.PagingArrears{
		ID:   userID,
		Page: page,
		Size: size,
	}
	arrearsServ, err := ap.arrears.GetPagedReadersArrears(ctx, in)
	if err != nil {
		log.Println("Arrear Client: Can't recieve arrears list")
		return nil, err
	}

	var arrears []Arrear

	for {
		recvArrear, err := arrearsServ.Recv()
		if err == io.EOF {
			log.Println("Arrear Client: All arrears recieved successfully")
			return arrears, nil
		} else if err != nil {
			log.Println("Arrear Client: Can't receive arrear")
			return nil, err
		}
		arrears = append(arrears,
			Arrear{
				ID:       recvArrear.GetID(),
				ReaderID: recvArrear.GetReaderID(),
				BookID:   recvArrear.GetBookID(),
				Start:    recvArrear.GetStart(),
				End:      recvArrear.GetEnd(),
			})
	}
}

func (ap *ArrearsPart) NewArrear(readerID int32, bookID int32) (*Arrear, error) {
	log.Println("Arrear Client: Registering new arrear for reader with ID", readerID, "and book ID", bookID)
	// header := metadata.New(map[string]string{"Authorization": secrets.AppKey + ":" + secrets.ArrearsSecret})
	// ctx := metadata.NewOutgoingContext(context.Background(), header)
	ctx := context.Background()

	newArrearReq := &protocol.NewArrear{
		ReaderID: readerID,
		BookID:   bookID,
	}

	arrear, err := ap.arrears.RegisterNewArrear(ctx, newArrearReq)
	if err != nil {
		log.Println("Arrear Client: Can't register new arrear")
		return nil, err
	}

	log.Println("Arrear Client: Arrear registered successfully")
	return &Arrear{
		ID:       arrear.GetID(),
		ReaderID: arrear.GetReaderID(),
		BookID:   arrear.GetBookID(),
		Start:    arrear.GetStart(),
		End:      arrear.GetEnd(),
	}, nil
}

func (ap *ArrearsPart) GetArrearByID(ID int32) (*Arrear, error) {
	log.Println("Arrear Client: Getting arrear with ID", ID)
	// header := metadata.New(map[string]string{"Authorization": secrets.AppKey + ":" + secrets.ArrearsSecret})
	// ctx := metadata.NewOutgoingContext(context.Background(), header)
	ctx := context.Background()

	arrearID := &protocol.SomeArrearsID{
		ID: ID,
	}

	arrear, err := ap.arrears.GetArrearByID(ctx, arrearID)
	if err != nil {
		log.Println("Arrear Client: Can't get arrear")
		return nil, err
	}

	log.Println("Arrear Client: Arrear received successfully")
	return &Arrear{
		ID:       arrear.GetID(),
		ReaderID: arrear.GetReaderID(),
		BookID:   arrear.GetBookID(),
		Start:    arrear.GetStart(),
		End:      arrear.GetEnd(),
	}, nil
}

func (ap *ArrearsPart) CloseArrearByID(ID int32) error {
	log.Println("Arrear Client: Close register with ID", ID)
	// header := metadata.New(map[string]string{"Authorization": secrets.AppKey + ":" + secrets.ArrearsSecret})
	// ctx := metadata.NewOutgoingContext(context.Background(), header)
	ctx := context.Background()
	req := &protocol.SomeArrearsID{ID: ID}
	_, err := ap.arrears.DeleteArrearByID(ctx, req)
	if err != nil {
		log.Println("Arrear Client: Can't close arrear")
	} else {
		log.Println("Arrear Client: Arrear closed succesfully")
	}
	return err
}

func AuthA(books protocol.ArrearsClient) (string, error) {
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

func WaithAndRefreshA(bp *ArrearsPart) {
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

	newToken, _ := AuthA(bp.arrears)
	bp.token = newToken
	go WaithAndRefreshA(bp)
}
