package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ersa97/new-grpc/server/data"
	"github.com/ersa97/new-grpc/server/service"
	"google.golang.org/grpc"
)

func main() {
	conn, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatal("make sure GRPC server is running")
	}

	s := service.Server{}

	grpcServer := grpc.NewServer()
	data.RegisterAllServer(grpcServer, &s)

	fmt.Println("GRPC Server running at http://localhost:8000")

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal("Failed to connect")
	}

}
