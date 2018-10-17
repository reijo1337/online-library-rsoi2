package main

import (
	"log"
	"net"

	"github.com/reijo1337/online-library-rsoi2/arrears-service/protocol"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalln("can't listet port", err)
	}

	serv, err := Server()
	if err != nil {
		log.Fatalln("can't  start server", err)
	}

	server := grpc.NewServer()

	protocol.RegisterArrearsServer(server, serv)
	log.Println("starting server at :8083")
	server.Serve(lis)
}
