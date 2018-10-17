package main

import (
	"context"
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
		addr = "0.0.0.0"
	}

	grpcConn, err := grpc.Dial(addr + ":8082")
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

	nothing := &protocol.Nothing{Dummy: true}

	readersServer, err := rp.readers.GetReadersList(ctx, nothing)
	if err != nil {
		return nil, err
	}

	var readers []Reader

	for {
		recvReader, err := readersServer.Recv()
		readers = append(readers,
			Reader{
				ID:   recvReader.GetID(),
				Name: recvReader.GetName(),
			})
		if err == io.EOF {
			return readers, nil
		} else if err != nil {
			return nil, err
		}
	}
}

func (rp *ReadersPart) getReaderByName(name string) (Reader, error) {
	ctx := context.Background()

	getReaderRequest := &protocol.ReaderName{Name: name}

	reader, err := rp.readers.GetReaderByName(ctx, getReaderRequest)

	if err != nil {
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
