package client

import (
	"client-streaming/pb"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func Run() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	client := pb.NewTemperatureServiceClient(dial)

	stream, err := client.RecordTemperatura(context.Background())
	if err != nil {
		panic(err)
	}

	temperatures := []float32{19.5, 20.0, 21.0, 22.0, 23.0, 24.0, 25.0}

	for _, temperature := range temperatures {
		fmt.Printf("Sending temperature: %.2f\n", temperature)

		if err := stream.Send(&pb.TemperatureRequest{Temperature: temperature}); err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)
	}

	recv, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Average temperature: %.2f\n", recv.GetAverageTemperature())

}
