package client

import (
	"context"
	"fmt"
	"io"
	"server-streaming/pb"

	"google.golang.org/grpc"
)

func Run() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	client := pb.NewStatusClient(dial)

	stream, err := client.StreamStatus(context.Background(), &pb.StreamRequest{TaskId: "123"})
	if err != nil {
		panic(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Printf("Received status: %s, progress: %d%%\n", res.GetMessage(), res.GetProgress())
	}

}
