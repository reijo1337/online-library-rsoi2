package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/reijo1337/online-library-rsoi2/arrears-service/protocol"
	"google.golang.org/grpc/metadata"
)

type ArrearServer struct {
	db *Database
}

func Server() (*ArrearServer, error) {
	log.Println("Set up arrear service...")
	db, err := SetUpDatabase()
	if err != nil {
		return nil, err
	}
	return &ArrearServer{db: db}, nil
}

func isAuthorized(ctx context.Context) bool {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ok
	}

	tokenString := headers["authorization"][0]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(Secret), nil
	})

	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	} else {
		return false
	}
}

func (as *ArrearServer) GetPagedReadersArrears(in *protocol.PagingArrears, p protocol.Arrears_GetPagedReadersArrearsServer) error {
	log.Println("Server: New request for arrears with pagging. User ID:", in.ID, ", page:", in.Page, ", page size:", in.Size)
	arrears, err := as.db.GetArrearsPaggin(in.ID, in.Size, in.Page)
	if err != nil {
		log.Println("Server: Can't process this request: ", err.Error())
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
			log.Println("Server: Can't send arrear: ", err.Error())
			return err
		}
	}
	log.Println("Server: Request processed successfully")
	return nil
}

func (as *ArrearServer) RegisterNewArrear(ctx context.Context, in *protocol.NewArrear) (*protocol.Arrear, error) {
	log.Println("Server: New request for arrear registration with", in.GetBookID(), "book ID and", in.GetReaderID(), "reader ID")
	startTime := time.Now()
	endTime := startTime.AddDate(0, 1, 0)

	start := parseDate(startTime)
	end := parseDate(endTime)

	log.Println("Server: arrear period [", start, ";", end, "]")

	arrear, err := as.db.InsertNewArrear(in.GetReaderID(), in.GetBookID(), start, end)
	if err != nil {
		log.Println("Server: Can't make new arrear:", err.Error())
		return nil, err
	}
	log.Println("Server: Request processed successfully")
	return &protocol.Arrear{
		ID:       arrear.ID,
		ReaderID: arrear.readerID,
		BookID:   arrear.bookID,
		Start:    arrear.start,
		End:      arrear.end,
	}, nil
}

func (as *ArrearServer) GetArrearByID(ctx context.Context, in *protocol.SomeArrearsID) (*protocol.Arrear, error) {
	log.Println("Server: New request for arrear with", in.GetID(), "ID")
	arrear, err := as.db.GetArrearByID(in.GetID())
	if err != nil {
		log.Println("Server: Can't getting arrear:", err.Error())
		return nil, err
	}
	log.Println("Server: Request processed successfully")
	return &protocol.Arrear{
		ID:       arrear.ID,
		ReaderID: arrear.readerID,
		BookID:   arrear.bookID,
		Start:    arrear.start,
		End:      arrear.end,
	}, nil
}

func (as *ArrearServer) DeleteArrearByID(ctx context.Context, in *protocol.SomeArrearsID) (*protocol.NothingArrear, error) {
	log.Println("Server: New request for deleting arrear with", in.GetID(), "ID")
	err := as.db.CloseArrayByID(in.GetID())

	if err != nil {
		log.Println("Server: Can't delete this arrear")
	} else {
		log.Println("Server: Request processed successfully")
	}
	return &protocol.NothingArrear{Dummy: true}, err
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

func (s *ArrearServer) Auth(ctx context.Context, in *protocol.AuthRequest) (*protocol.SomeString, error) {
	log.Println("Server: New authorization")
	if ContainsKey(in.AppKey) && (in.AppSecret == Secret) {
		token, err := genToken(in.AppKey)
		return &protocol.SomeString{String_: token}, err
	}
	return nil, errors.New("Unauthorized")
}

func genToken(login string) (string, error) {
	log.Println("Server: Generating token")
	hmacSampleSecret := []byte(Secret)
	AccessTokenExp := time.Now().Add(time.Minute * 30).Unix()
	log.Println("Server: Gen access token")
	accesToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "BookService",
		"exp": AccessTokenExp,
		"aud": login,
	})

	log.Println("Server: Signing access token", accesToken, hmacSampleSecret)
	accessTokenString, err := accesToken.SignedString(hmacSampleSecret)
	if err != nil {
		log.Println("Server: Can't authorize: ", err.Error())
		return "", err
	}

	return accessTokenString, nil
}
