package main

import (
	"context"
	"strconv"
	"time"

	"github.com/reijo1337/online-library-rsoi2/arrears-service/protocol"
)

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

func (as *ArrearServer) RegisterNewArrear(ctx context.Context, in *protocol.NewArrear) (*protocol.Arrear, error) {
	startTime := time.Now()
	endTime := startTime.AddDate(0, 1, 0)

	start := parseDate(startTime)
	end := parseDate(endTime)

	arrear, err := as.db.InsertNewArrear(in.GetReaderID(), in.GetBookID(), start, end)
	if err != nil {
		return nil, err
	}
	return &protocol.Arrear{
		ID:       arrear.ID,
		ReaderID: arrear.readerID,
		BookID:   arrear.bookID,
		Start:    arrear.start,
		End:      arrear.end,
	}, nil
}

func parseDate(t time.Time) string {
	year := strconv.Itoa(t.Year())
	month := strconv.Itoa(int(t.Month()))
	if len(month) == 1 {
		month = "0" + month
	}
	day := strconv.Itoa(t.Day())
	if len(day) == 1 {
		day = "0" + day
	}

	return year + month + day
}
