package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"
	"unary/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func Run() {
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	dial, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	userClient := pb.NewUserClient(dial)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	loginResp, err := userClient.Login(ctx, &pb.LoginRequest{Username: "user", Password: "password"})
	if err != nil {
		log.Fatalf("Falha ao fazer login: %v", err)
	}

	token := loginResp.Token
	log.Printf("Token JWT recebido: %s", token)

	md := metadata.New(map[string]string{"authorization": "Bearer " + token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	user, err := userClient.AddUser(ctx, &pb.AddUserRequest{
		Id:   "1",
		Name: "John",
		Age:  30,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("First user created: %v\n", user)

	user, err = userClient.AddUser(ctx, &pb.AddUserRequest{
		Id:   "2",
		Name: "Paul",
		Age:  42,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Second user created: %v\n", user)

	getUser, err := userClient.GetUser(ctx, &pb.GetUserRequest{Id: "2"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("User returned from GetUser method: %v\n", getUser)

	time.Sleep(1 * time.Second)

}
