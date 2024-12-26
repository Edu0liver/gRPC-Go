package client

import (
	"context"
	"fmt"
	"unary/pb"

	"google.golang.org/grpc"
)

func Run() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	userClient := pb.NewUserClient(dial)

	user, err := userClient.AddUser(context.Background(), &pb.AddUserRequest{
		Id:   "1",
		Name: "John",
		Age:  30,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("First User created: %v\n", user)

	user, err = userClient.AddUser(context.Background(), &pb.AddUserRequest{
		Id:   "2",
		Name: "Paul",
		Age:  42,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Second User created: %v\n", user)

	getUser, err := userClient.GetUser(context.Background(), &pb.GetUserRequest{
		Id: "2",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("User returned from GetUser method: %v\n", getUser)
}
