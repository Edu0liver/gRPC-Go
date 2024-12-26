package server

import (
	"bidirectional-streaming/pb"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

type StockServiceServer struct {
	pb.UnimplementedStockServiceServer
}

func (s *StockServiceServer) StreamStockPrices(stream grpc.BidiStreamingServer[pb.StockRequest, pb.StockResponse]) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}

			return err
		}

		symbol := req.GetSymbol()
		fmt.Printf("Received symbol: %s\n", symbol)

		go func(symbol string) {
			for i := 0; i < 50; i++ {
				price := rand.Float32() * 100

				if err := stream.Send(&pb.StockResponse{Symbol: symbol, Price: price}); err != nil {
					panic(fmt.Sprintf("Error sending price: %s\n", err))
				}
			}
		}(symbol)

		time.Sleep(1 * time.Second)
	}
}

func Run() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStockServiceServer(grpcServer, &StockServiceServer{})
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
