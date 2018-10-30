package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/reijo1337/online-library-rsoi2/arrears-service/protocol"
	"google.golang.org/grpc"
)

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

func (ap *ArrearsPart) getArrearsPaging(userID int32, page int32, size int32) ([]Arrear, error) {
	log.Println("Arrear Client: Getting arrears with pagging. User ID:", userID, ", page:", page, ", page size:", size)
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
				readerID: recvArrear.GetReaderID(),
				bookID:   recvArrear.GetBookID(),
				start:    recvArrear.GetStart(),
				end:      recvArrear.GetEnd(),
			})
	}
}

func (ap *ArrearsPart) newArrear(readerID int32, bookID int32) (*Arrear, error) {
	log.Println("Arrear Client: Registering new arrear for reader with ID", readerID, "and book ID", bookID)
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
		readerID: arrear.GetReaderID(),
		bookID:   arrear.GetBookID(),
		start:    arrear.GetStart(),
		end:      arrear.GetEnd(),
	}, nil
}

func (ap *ArrearsPart) getArrearByID(ID int32) (*Arrear, error) {
	log.Println("Arrear Client: Getting arrear with ID", ID)
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
		readerID: arrear.GetReaderID(),
		bookID:   arrear.GetBookID(),
		start:    arrear.GetStart(),
		end:      arrear.GetEnd(),
	}, nil
}

func (ap *ArrearsPart) closeArrearByID(ID int32) error {
	log.Println("Arrear Client: Close register with ID", ID)
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
