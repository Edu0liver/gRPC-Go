package server

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"
	"time"
	"unary/pb"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatalf("Erro ao carregar certificados TLS: %v", err)
	}

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(ChainUnaryInterceptors(authInterceptor, rateLimitInterceptor, logInterceptor, timeoutInterceptor)))

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

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Username == "user" && req.Password == "password" {
		token, err := generateJWT(req.Username)
		if err != nil {
			return nil, err
		}

		return &pb.LoginResponse{Token: token}, nil
	}

	return nil, errors.New("usuário ou senha inválidos")
}

func generateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString(jwtKey)
}
