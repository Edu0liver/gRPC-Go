package server

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var jwtKey = []byte("t798VT68v6798v6v89G679VG678GSY9D")

func validateJWT(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	}
	return "", errors.New("token inválido")
}

func authInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	if info.FullMethod == "/pb.User/Login" {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadados ausentes")
	}

	token := md["authorization"]
	if len(token) == 0 {
		return nil, errors.New("token ausente")
	}

	username, err := validateJWT(token[0][7:])
	if err != nil {
		return nil, err
	}

	newCtx := context.WithValue(ctx, "username", username)

	return handler(newCtx, req)
}

func ChainUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		// Função handler que chamará o próximo interceptador da cadeia
		currentHandler := handler

		for i := len(interceptors) - 1; i >= 0; i-- {
			interceptor := interceptors[i]
			next := currentHandler
			currentHandler = func(currentCtx context.Context, currentReq any) (any, error) {
				return interceptor(currentCtx, currentReq, info, next)
			}
		}

		return currentHandler(ctx, req)
	}
}

var limiter = rate.NewLimiter(1, 5)
var mu sync.Mutex

func rateLimitInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	mu.Lock()
	defer mu.Unlock()
	if !limiter.Allow() {
		return nil, status.Error(codes.ResourceExhausted, "limite de requisições excedido")
	}
	return handler(ctx, req)
}

func logInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	log.Printf("METHOD: %s", info.FullMethod)

	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("ERROR - METHOD: %s, Error: %s", info.FullMethod, err.Error())
	}

	return resp, err
}

func timeoutInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return handler(ctx, req)
}
