package main

import (
	"fmt"
	"log"

	"github.com/ersa97/new-grpc/server/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Users struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Pass     string `json:"pass"`
}

var (
	grpcClient   data.AuthenticationClient
	log_database = []Users{
		{
			Id:       1,
			Username: "ersa",
			Pass:     "1234567890",
		},
		{
			Id:       2,
			Username: "Adinda",
			Pass:     "0987654321",
		},
	}
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("GRPC Client failed to connect to GRPC Server")
	}
	defer conn.Close()
	grpcClient = data.NewAuthenticationClient(conn)

	fmt.Println("CLient is Running")

	// ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	// defer cancel()

	// r1, err := grpcClient.GetUsers(ctx, &data.GetUsersRequest{})
	// if err != nil {
	// 	log.Fatalf("get users %s", err)
	// }

	// _ := r1.User

	fmt.Println("todos:", log_database)

}
