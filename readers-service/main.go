package main

import (
	"log"
	"net"

	"github.com/reijo1337/online-library-rsoi2/readers-service/protocol"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Println("can't listet port", err)
	}

	serv, err := Server()
	if err != nil {
		log.Println("can't  start server", err)
	}

	server := grpc.NewServer()

	protocol.RegisterReadersServer(server, serv)
	log.Println("starting server at :8082")
	server.Serve(lis)
}
