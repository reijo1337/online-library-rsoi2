package main

import "github.com/reijo1337/online-library-rsoi2/arrears-service/protocol"

type ArrearServer struct {
	db *Database
}

func Server() (*ArrearServer, error) {
	db, err := SetUpDatabase()
	if err != nil {
		return nil, err
	}
	return &ArrearServer{db: db}, nil
}

func (as *ArrearServer) GetPagedReadersArrears(in *protocol.PagingArrears, p protocol.Arrears_GetPagedReadersArrearsServer) error {
	arrears, err := as.db.GetArrearsPaggin(in.ID, in.Size, in.Page)
	if err != nil {
		return err
	}
	for _, arrear := range arrears {
		ar := &protocol.Arrear{
			ID:       arrear.ID,
			ReaderID: arrear.readerID,
			BookID:   arrear.bookID,
			Start:    arrear.start,
			End:      arrear.end,
		}
		if err := p.Send(ar); err != nil {
			return err
		}
	}
	return nil
}
