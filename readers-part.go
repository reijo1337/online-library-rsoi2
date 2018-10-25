package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"google.golang.org/grpc"

	"github.com/reijo1337/online-library-rsoi2/readers-service/protocol"
)

type ReadersPart struct {
	conn    *grpc.ClientConn
	readers protocol.ReadersClient
}

func NewReadersPart() (*ReadersPart, error) {
	addr := os.Getenv("READERSADDR")
	if addr == "" {
		addr = "127.0.0.1"
	}

	grpcConn, err := grpc.Dial(
		addr+":8082",
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	readers := protocol.NewReadersClient(grpcConn)
	return &ReadersPart{
		conn:    grpcConn,
		readers: readers,
	}, nil
}

func (rp *ReadersPart) getAllReaders() ([]Reader, error) {
	ctx := context.Background()

	nothing := &protocol.NothingReaders{Dummy: true}

	readersServer, err := rp.readers.GetReadersList(ctx, nothing)
	if err != nil {
		return nil, err
	}

	var readers []Reader

	for {
		recvReader, err := readersServer.Recv()
		if err == io.EOF {
			return readers, nil
		} else if err != nil {
			return nil, err
		}
		readers = append(readers,
			Reader{
				ID:   recvReader.GetID(),
				Name: recvReader.GetName(),
			})
	}
}

func (rp *ReadersPart) getReaderByName(name string) (Reader, error) {
	ctx := context.Background()

	getReaderRequest := &protocol.ReaderName{Name: name}

	reader, err := rp.readers.GetReaderByName(ctx, getReaderRequest)

	if err != nil {
		fmt.Println(err.Error())
		return Reader{}, err
	}

	return Reader{ID: reader.GetID(), Name: reader.GetName()}, nil
}

func (rp *ReadersPart) getReaderByID(ID int32) (Reader, error) {
	ctx := context.Background()

	getReaderRequest := &protocol.ReaderID{ID: ID}

	reader, err := rp.readers.GetReaderByID(ctx, getReaderRequest)

	if err != nil {
		return Reader{}, err
	}

	return Reader{ID: reader.GetID(), Name: reader.GetName()}, nil
}

func (rp *ReadersPart) registerReader(name string) (*Reader, error) {
	ctx := context.Background()
	readerNameReq := &protocol.ReaderName{Name: name}

	reader, err := rp.readers.RegisterReader(ctx, readerNameReq)
	if err != nil {
		return nil, err
	}
	return &Reader{ID: reader.GetID(), Name: reader.GetName()}, nil
}
