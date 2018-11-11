package main

import (
	"context"
	"log"

	"github.com/reijo1337/online-library-rsoi2/readers-service/protocol"
)

type ReadersServer struct {
	db *Database
}

func Server() (*ReadersServer, error) {
	log.Println("Set up reader service...")
	db, err := SetUpDatabase()
	if err != nil {
		return nil, err
	}
	return &ReadersServer{db: db}, nil
}

func (s *ReadersServer) RegisterReader(ctx context.Context, user *protocol.ReaderName) (*protocol.Reader, error) {
	log.Println("Server: New request for registering new reader with name", user.GetName())
	reader, err := s.db.addReader(user.GetName())
	if err != nil {
		log.Println("Server: Can't process this request:", err.Error())
		return nil, err
	}
	log.Println("Server: Request processed successfully")
	return &protocol.Reader{ID: reader.ID, Name: reader.Name}, nil
}

func (s *ReadersServer) GetReadersList(in *protocol.NothingReaders, p protocol.Readers_GetReadersListServer) error {
	log.Println("Server: New request for readers list")
	readers, err := s.readerList()
	if err != nil {
		log.Println("Server: Can't process this request: ", err.Error())
		return err
	}

	for _, reader := range readers {
		if err := p.Send(reader); err != nil {
			log.Println("Server: Can't send reader: ", err.Error())
			return err
		}
	}

	log.Println("Server: Request processed successfully")
	return nil
}

func (s *ReadersServer) readerList() ([]*protocol.Reader, error) {
	readers, err := s.db.getReadersList()
	if err != nil {
		return nil, err
	}

	ret := make([]*protocol.Reader, len(readers))
	for _, reader := range readers {
		ret = append(ret, &protocol.Reader{ID: reader.ID, Name: reader.Name})
	}

	return ret, nil
}

func (s *ReadersServer) GetReaderByName(ctx context.Context, user *protocol.ReaderName) (*protocol.Reader, error) {
	log.Println("Server: New request for reader with name", user.GetName())
	reader, err := s.db.getReaderByName(user.GetName())
	if err != nil {
		log.Println("Server: Can't process this request: ", err.Error())
		return nil, err
	}
	log.Println("Server: Request processed successfully")
	return &protocol.Reader{ID: reader.ID, Name: reader.Name}, nil
}

func (s *ReadersServer) GetReaderByID(ctx context.Context, user *protocol.ReaderID) (*protocol.Reader, error) {
	log.Println("Server: New request for reader with ID", user.GetID())
	reader, err := s.db.getReaderByID(user.GetID())
	if err != nil {
		log.Println("Server: Can't process this request: ", err.Error())
		return nil, err
	}
	log.Println("Server: Request processed successfully")
	return &protocol.Reader{ID: reader.ID, Name: reader.Name}, nil
}
