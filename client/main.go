package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ersa97/new-grpc/client/data"
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
	fmt.Println(r1)
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
			Password: []byte("mustika"),
		},
	})
	if err != nil {
		log.Fatalf("register %s", err)
	}

	fmt.Println(r1.Message)
	getAll()
}

func Login() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	r1, err := grpcClient.Login(ctx, &data.LoginRequest{
		Email:    "mustikadk1999@gmail.com",
		Password: []byte("mustika"),
	})
	if err != nil {
		log.Fatalf("login %s", err)
	}
	fmt.Println(r1.AccessToken)

	var opt string
	fmt.Println("Welcome mustikadk1999@gmail.com \nPlease pick a menu to start\n1. Add user Users\n2. Delete User\n3. Update")
	fmt.Scanf("%s", &opt)
	fmt.Println(opt)
	if opt == "1" {
		Add(r1.AccessToken)
	}
}

func Add(token string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	r1, err := grpcClient.AddUser(ctx, &data.AddUserRequest{
		AccessToken: token,
		User: &data.User{
			Name:     "Safrizal",
			Email:    "safrizal99@gmail.com",
			Password: []byte("safrizal99"),
		},
	})
	if err != nil {
		log.Fatalf("add user %s", err)
	}

	fmt.Println("Users : ")
	fmt.Printf("ID\t: %v\nName\t: %v\nEmail\t: %v\nPassword: %v\n\n", r1.User.Id, r1.User.Name, r1.User.Email, r1.User.Password)

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
	fmt.Println("Welcome to User Configuration \n please pick a menu to start\n1. get All Users\n2. Register\n3. Login")
	fmt.Scanf("%s", &menu)

	switch menu {
	case "1":
		getAll()
	case "2":
		Register()
	case "3":
		Login()

	}
}
