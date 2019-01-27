package main

import (
	"log"
	"net"
	"time"

	"github.com/reijo1337/online-library-rsoi2/readers-service/protocol"
	"google.golang.org/grpc"
)

func main() {
	var (
		lis  net.Listener
		err  error
		serv *ReadersServer
	)
	for {
		lis, err = net.Listen("tcp", ":8082")
		if err != nil {
			log.Println("can't listet port", err)
			log.Println("Retry after 5 sec")
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	for {
		serv, err = Server()
		if err != nil {
			log.Println("can't  start server", err)
			log.Println("Retry after 5 sec")
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	server := grpc.NewServer()

	protocol.RegisterReadersServer(server, serv)
	log.Println("starting server at :8082")
	server.Serve(lis)
}
