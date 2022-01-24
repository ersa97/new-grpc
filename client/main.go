package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ersa97/new-grpc/server/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcClient data.AuthenticationClient
)

func getAll() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	r1, err := grpcClient.GetUsers(ctx, &data.GetUsersRequest{})
	if err != nil {
		log.Fatalf("get users %s", err)
	}

	fmt.Println("Users : ")
	for _, v := range r1.User {
		fmt.Printf("ID\t: %v\nName\t: %v\nEmail\t: %v\nPassword: %v\n\n", v.Id, v.Name, v.Email, v.Password)
	}
}

func Register() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	r1, err := grpcClient.RegisterUser(ctx, &data.RegisterRequest{
		User: &data.User{
			Name:     "Mustika",
			Email:    "mustikadk1999@gmail.com",
			Password: "mustika",
		},
	})
	if err != nil {
		log.Fatalf("register %s", err)
	}

	fmt.Println(r1.Message)
}

func main() {
	conn, err := grpc.Dial("0.0.0.0:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("GRPC Client failed to connect to GRPC Server")
	}
	defer conn.Close()
	grpcClient = data.NewAuthenticationClient(conn)

	fmt.Println("CLient is Running")

	var menu string
	fmt.Println("Welcome to User Configuration \n please pick a menu to start\n1. get All Users\n2. Register")
	fmt.Scanf("%s", &menu)

	switch menu {
	case "1":
		getAll()
	case "2":
		Register()
		getAll()
	}
}
