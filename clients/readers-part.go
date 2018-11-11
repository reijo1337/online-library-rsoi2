package clients

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/reijo1337/online-library-rsoi2/readers-service/protocol"
)

type ReadersPartInterface interface {
	GetAllReaders() ([]Reader, error)
	GetReaderByName(name string) (Reader, error)
	GetReaderByID(ID int32) (Reader, error)
	RegisterReader(name string) (*Reader, error)
}

type ReadersPart struct {
	conn    *grpc.ClientConn
	readers protocol.ReadersClient
}

func NewReadersPart() (*ReadersPart, error) {
	log.Println("Readers Client: Connecting to readers service...")
	addr := os.Getenv("READERSADDR")
	if addr == "" {
		addr = "127.0.0.1"
	}

	log.Println("Readers Client: readers service addres:", addr+":8082")
	grpcConn, err := grpc.Dial(
		addr+":8082",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Println("Readers Client: Can't connect to remote service")
		return nil, err
	}

	readers := protocol.NewReadersClient(grpcConn)
	log.Println("Readers Client: success!")
	return &ReadersPart{
		conn:    grpcConn,
		readers: readers,
	}, nil
}

func (rp *ReadersPart) GetAllReaders() ([]Reader, error) {
	log.Println("Readers Client: Getting readers list")
	ctx := context.Background()

	nothing := &protocol.NothingReaders{Dummy: true}

	readersServer, err := rp.readers.GetReadersList(ctx, nothing)
	if err != nil {
		log.Println("Readers Client: Can't recieve readers list")
		return nil, err
	}

	var readers []Reader

	for {
		recvReader, err := readersServer.Recv()
		if err == io.EOF {
			log.Println("Readers Client: All readers received succesfully")
			return readers, nil
		} else if err != nil {
			log.Println("Readers Client: Can't recieve one reader")
			return nil, err
		}
		readers = append(readers,
			Reader{
				ID:   recvReader.GetID(),
				Name: recvReader.GetName(),
			})
	}
}

func (rp *ReadersPart) GetReaderByName(name string) (Reader, error) {
	log.Println("Readers Client: Getting reader by name", name)
	ctx := context.Background()

	getReaderRequest := &protocol.ReaderName{Name: name}

	reader, err := rp.readers.GetReaderByName(ctx, getReaderRequest)

	if err != nil {
		fmt.Println("Readers Client: Can't recieve reader")
		return Reader{}, err
	}

	log.Println("Readers Client: Reader recieved succesfully")
	return Reader{ID: reader.GetID(), Name: reader.GetName()}, nil
}

func (rp *ReadersPart) GetReaderByID(ID int32) (Reader, error) {
	log.Println("Readers Client: Getting reader by ID", ID)
	ctx := context.Background()

	getReaderRequest := &protocol.ReaderID{ID: ID}

	reader, err := rp.readers.GetReaderByID(ctx, getReaderRequest)

	if err != nil {
		log.Println("Readers Client: Can't recieve reader")
		return Reader{}, err
	}

	log.Println("Readers Client: Reader recieved succesfully")
	return Reader{ID: reader.GetID(), Name: reader.GetName()}, nil
}

func (rp *ReadersPart) RegisterReader(name string) (*Reader, error) {
	log.Println("Readers Client: Registration of new reader")
	ctx := context.Background()
	readerNameReq := &protocol.ReaderName{Name: name}

	reader, err := rp.readers.RegisterReader(ctx, readerNameReq)
	if err != nil {
		log.Println("Readers Client: Can't register reader")
		return nil, err
	}
	log.Println("Readers Client: Reader registered succesfully")
	return &Reader{ID: reader.GetID(), Name: reader.GetName()}, nil
}
