package main

import (
	"context"
	"io"
	"os"

	"github.com/reijo1337/online-library-rsoi2/arrears-service/protocol"
	"google.golang.org/grpc"
)

type ArrearsPart struct {
	conn    *grpc.ClientConn
	arrears protocol.ArrearsClient
}

func NewArrearsPart() (*ArrearsPart, error) {
	addr := os.Getenv("ARREARSADDR")
	if addr == "" {
		addr = "0.0.0.0"
	}

	grpcConn, err := grpc.Dial(
		addr+":8083",
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, err
	}

	arrears := protocol.NewArrearsClient(grpcConn)
	return &ArrearsPart{
		conn:    grpcConn,
		arrears: arrears,
	}, nil
}

func (ap *ArrearsPart) getArrearsPaging(userID int32, page int32, size int32) ([]Arrear, error) {
	ctx := context.Background()
	in := &protocol.PagingArrears{
		ID:   userID,
		Page: page,
		Size: size,
	}
	arrearsServ, err := ap.arrears.GetPagedReadersArrears(ctx, in)
	if err != nil {
		return nil, err
	}

	var arrears []Arrear

	for {
		recvArrear, err := arrearsServ.Recv()
		if err == io.EOF {
			return arrears, nil
		} else if err != nil {
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
	ctx := context.Background()

	newArrearReq := &protocol.NewArrear{
		ReaderID: readerID,
		BookID:   bookID,
	}

	arrear, err := ap.arrears.RegisterNewArrear(ctx, newArrearReq)
	if err != nil {
		return nil, err
	}

	return &Arrear{
		ID:       arrear.GetID(),
		readerID: arrear.GetReaderID(),
		bookID:   arrear.GetBookID(),
		start:    arrear.GetStart(),
		end:      arrear.GetEnd(),
	}, nil
}
