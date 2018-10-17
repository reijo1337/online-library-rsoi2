package main

import (
	"context"

	"github.com/reijo1337/online-library-rsoi2/readers-service/protocol"
)

type ReadersServer struct {
	db *Database
}

func Server() (*ReadersServer, error) {
	db, err := SetUpDatabase()
	if err != nil {
		return nil, err
	}
	return &ReadersServer{db: db}, nil
}

func (s *ReadersServer) RegisterReader(ctx context.Context, user *protocol.ReaderName) (*protocol.Reader, error) {
	reader, err := s.db.addReader(user.GetName())
	if err != nil {
		return nil, err
	}
	return &protocol.Reader{ID: reader.ID, Name: reader.Name}, nil
}

func (s *ReadersServer) GetReadersList(in *protocol.Nothing, p protocol.Readers_GetReadersListServer) error {
	readers, err := s.readerList()
	if err != nil {
		return err
	}

	for _, reader := range readers {
		if err := p.Send(reader); err != nil {
			return err
		}
	}

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
	reader, err := s.db.getReaderByName(user.GetName())
	if err != nil {
		return nil, err
	}
	return &protocol.Reader{ID: reader.ID, Name: reader.Name}, nil
}

func (s *ReadersServer) GetReaderByID(ctx context.Context, user *protocol.ReaderID) (*protocol.Reader, error) {
	reader, err := s.db.getReaderByID(user.GetID())
	if err != nil {
		return nil, err
	}
	return &protocol.Reader{ID: reader.ID, Name: reader.Name}, nil
}
