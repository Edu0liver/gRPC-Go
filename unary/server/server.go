package server

import (
	"context"
	"errors"
	"net"
	"sync"
	"unary/pb"

	"google.golang.org/grpc"
)

type User struct {
	ID   string
	Name string
	Age  int32
}

type UserService struct {
	pb.UnimplementedUserServer

	users map[string]*User
	mu    sync.Mutex
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*User),
	}
}

func Run() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterUserServer(s, NewUserService())

	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}

func (us *UserService) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	user := &User{
		ID:   req.Id,
		Name: req.Name,
		Age:  req.Age,
	}

	us.users[user.ID] = user

	return &pb.AddUserResponse{
		Id:   user.ID,
		Age:  user.Age,
		Name: user.Name,
	}, nil
}

func (us *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	user, ok := us.users[req.Id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &pb.GetUserResponse{
		Id:   user.ID,
		Age:  user.Age,
		Name: user.Name,
	}, nil
}
