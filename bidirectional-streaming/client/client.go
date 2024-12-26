package client

import (
	"bidirectional-streaming/pb"
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc"
)

func Run() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	client := pb.NewStockServiceClient(dial)

	stream, err := client.StreamStockPrices(context.Background())
	if err != nil {
		panic(err)
	}

	done := make(chan struct{})

	go func() {
		for {
			recv, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}

				break
			}

			fmt.Printf("Received stock price: %.2f\n", recv.GetPrice())
		}

		close(done)
	}()

	symbols := []string{"GOOGL", "AAPL", "MSFT", "AMZN", "FB", "TSLA"}

	for _, symbol := range symbols {
		fmt.Printf("Sending symbol: %s\n", symbol)

		if err := stream.Send(&pb.StockRequest{Symbol: symbol}); err != nil {
			panic(err)
		}

		time.Sleep(2 * time.Second)
	}

	<-done
}
