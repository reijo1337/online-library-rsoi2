package clients

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/reijo1337/online-library-rsoi2/arrears-service/protocol"
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
	return &ArrearsPart{
		conn:    grpcConn,
		arrears: arrears,
	}, nil
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
