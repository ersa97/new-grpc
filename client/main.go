package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:8000",grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!=nil{
		log.Fatal("GRPC Client failed to connect to GRPC Server")
	}

	defer conn.close()
}