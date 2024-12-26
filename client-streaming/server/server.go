package server

import (
	"client-streaming/pb"
	"net"

	"google.golang.org/grpc"
)

type TemperatureServer struct {
	pb.UnimplementedTemperatureServiceServer
}

func (s *TemperatureServer) RecordTemperatura(stream grpc.ClientStreamingServer[pb.TemperatureRequest, pb.TemperatureResponse]) error {
	var sum float32
	var count int32

	for {
		req, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				return stream.SendAndClose(&pb.TemperatureResponse{
					AverageTemperature: sum / float32(count),
				})
			}

			return err
		}

		sum += req.GetTemperature()
		count++
	}
}

func Run() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTemperatureServiceServer(grpcServer, &TemperatureServer{})
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
