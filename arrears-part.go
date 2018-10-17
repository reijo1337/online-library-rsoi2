package main

import (
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

	grpcConn, err := grpc.Dial(addr + ":8083")

	if err != nil {
		return nil, err
	}

	arrears := protocol.NewArrearsClient(grpcConn)
	return &ArrearsPart{
		conn:    grpcConn,
		arrears: arrears,
	}, nil
}
